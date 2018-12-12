var datatable;
var editor;

$( document ).ready(function() {

    $.get( "/restaurant/", function( data ) {
        if (data["error"]){
            $(".alert-danger strong").text(data["error"]);
            $(".alert-danger").removeClass("hide");
            if($('#restaurants').hasClass("dataTable")) {
                resetDataTable();
            }
        }
        else
        {
            $(".alert").addClass("hide");
            if($('#restaurants').hasClass("dataTable")) {
                refreshDataTable(data["restaurants"]);
            } else {
                datatable = initDataTable($('#restaurants'), data);
            }
        }
    });

    $("#restaurants-container").on("click", "#delete-btn", function(){
        let row = $(this).closest('tr')
        let resId = datatable.row(row).data()["id"];
        $.ajax({
            url: "/restaurant/" + resId,
            type: "DELETE",
            success: function(result) {
                if (result["success"]) {
                    datatable.row(row).remove().draw();
                }
            }
        });
        return false;
    });

    $("#restaurants-container").on("click", "#edit-btn", function(){
        let row = $(this).closest('tr')
        let rowData = datatable.row(row).data();
        $.each(rowData, function(key, value) {
            $("#" + key + "_input").val(value);
        });
        $('.modal').modal('show');
        return false;
    });

    $("form").submit(function() {
        let resId = $("#id_input").val();
        let formData = $('form').serializeArray();
        var jsonData = formData.reduce(function(map, obj) {
            map[obj.name] = obj.value;
            return map;
        }, {});
        $.ajax({
            url: "/restaurant/" + resId,
            type: "PUT",
            contentType: "json",
            data: JSON.stringify(jsonData),
            success: function(result) {
                if (result["restaurant"]){
                    $('.modal').modal('hide');
                    datatable.row('#' + resId).data(result["restaurant"]).draw();
                }
            },
            failure: function(result) {
                $('.modal').modal('hide');
            }
        });
        return false;
    });

    if (window["WebSocket"]) {
        try {
            tryConnectToReload("ws://localhost:12450/reload");
        }
        catch (ex) {
            console.error("rwetertrtert" +ex);
        }
    } else {
        console.log("Your browser does not support WebSockets.");
    }
});

function tryConnectToReload(address) {
    var conn = new WebSocket(address);
    conn.onclose = function () {
        setTimeout(function () {
            tryConnectToReload(address);
        }, 2000);
    };
    conn.onmessage = function (evt) {
        location.reload();
    };
}

function initDataTable(table, data) {
    let columns =  [
        {title: "ID", data: "id", width: "5%", className: "restaurant-id"},
        {title: "Name", data: "name", width: "15%"},
        {title: "Type", data: "type", width: "15%"},
        {title: "Phone", data: "phone", width: "15%"},
        {title: "Address", data: "address", width: "70%"},
        {title: "Edit/Delete", data: null, width: "20%",
            className: "center", defaultContent: '<a href="" id="edit-btn">Edit</a> / ' +
                '<a href="" id="delete-btn">Delete</a>'}
    ]
    let datatable = table.dataTable({
        columnDefs: [
            { "targets": [0,2,3,4,5], "searchable": false }
        ],
        paging: false,
        language: {
            searchPlaceholder: "Search by name",
            search: "",
        },
        searching: true,
        ordering: false,
        data: data["restaurants"],
        rowId: "id",
        columns: columns,
        "dom": '<"toolbar">frtip',
        'fnCreatedRow': function (nRow, aData, iDataIndex) {
            $(nRow).attr('resId', aData["id"]);
        },
    }).api();

    return datatable
}

function resetDataTable() {
    datatable.clear();
    datatable.draw();
}

function refreshDataTable(data) {
    datatable.clear();
    datatable.rows.add(data);
    datatable.draw();
}
