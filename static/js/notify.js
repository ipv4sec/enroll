
// lib
function logger(title, message, level) {
    type = 'pastel-danger';
    if (level === "info") {
        type = 'pastel-info'
    }
    $.notify({
        title: title,
        message: message
    },{
        type: type,
        delay: 8000,
        template: '<div data-notify="container" class="col-xs-11 col-sm-3 alert alert-{0}" role="alert">' +
            '<span data-notify="title">{1}</span>' +
            '<span data-notify="message">{2}</span>' +
            '</div>'
    });
}

