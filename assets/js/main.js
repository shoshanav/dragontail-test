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

    $("#sap-failed-invoices").on("click",".dt-mark-uploaded button", function(){
        let invoiceId = this.closest("tr").dataset.invoice_id
        let button = this;
        $.post( "/invoices/mark_as_uploaded/" + invoiceId, function( data ) {
            if (data["success"]){
                $(".alert-info strong").text(data["success"]);
                $(".alert-info").removeClass("hide");
                $(button).addClass("btn-success").html("MARKED");
            }
        });
    });

    $( "#search-by-name-form" ).submit(function() {
        $( ':button[type="submit"]' ).prop('disabled', true);
        $.get( "/restaurant/search?restaurantName=" + $("#res-name-input").val(), function( data ) {
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
                    getGeoCode(data);
                    datatable = initDataTable($('#restaurants'), data);
                }
            }
            $(':button[type="submit"]').prop('disabled', false);
        });
        return false;
    });

    // Edit record
    $('#restaurants').on('click', 'a.editor_edit', function (e) {
        e.preventDefault();

        editor.edit( $(this).closest('tr'), {
            title: 'Edit record',
            buttons: 'Update'
        } );
    } );

    // Delete a record
    $('#restaurants').on('click', 'a.editor_remove', function (e) {
        e.preventDefault();

        editor.remove( $(this).closest('tr'), {
            title: 'Delete record',
            message: 'Are you sure you wish to remove this record?',
            buttons: 'Delete'
        } );
    } );

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
        {title: "Edit/Delete", data: null, width: "20%", className: "center", defaultContent: '<a href="" class="editor_edit">Edit</a> / <a href="" class="editor_remove">Delete</a>'}
    ]
    let datatable = table.dataTable({
        columnDefs: [
            { "targets": [0,2,3], "searchable": false }
        ],
        paging: false,
        language: {
            searchPlaceholder: "Search by name",
            search: "",
        },
        searching: true,
        ordering: false,
        data: data["restaurants"],
        columns: columns,
        "dom": '<"toolbar">frtip',
        'fnCreatedRow': function (nRow, aData, iDataIndex) {
            $(nRow).attr('data-restaurant-id', aData["ID"]);
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
