var Notice = (function () {
    function error(msg) {
        showNotice(msg, 'danger')
    }

    function warning(msg) {
        showNotice(msg, 'warning')
    }

    function notice(msg) {
        showNotice(msg, 'info')
    }

    function success(msg) {
        showNotice(msg, 'success')
    }

    function showNotice(msg, type) {
        clear();
        $(`.alert-${type} strong`).text(msg);
        $(`.alert-${type}`).removeClass('hide');
    }

    function clear() {
        $('.alert').addClass('hide');
    }

    return {
        error: error,
        warning: warning,
        notice: notice,
        success: success,
        clear: clear
    }
})();

var DateFormat = (function () {
    function month(monthDate) {
        var localTime = moment.utc(monthDate)._d;
        return moment(localTime).format('MM/YY');
    }

    return {
        month: month
    }

})();

$(document).ajaxSuccess(function (event, xhr, settings) {
    if (xhr.responseJSON['warning']) {
        Notice.warning(xhr.responseJSON['warning'])
    } else if (xhr.responseJSON['success']) {
        Notice.success(xhr.responseJSON['success'])
    } else if (xhr.responseJSON['error']) {
        Notice.error(xhr.responseJSON['error'])
    } else if (xhr.responseJSON['info']) {
        Notice.notice(xhr.responseJSON['info'])
    } else {
        Notice.clear();
    }
});