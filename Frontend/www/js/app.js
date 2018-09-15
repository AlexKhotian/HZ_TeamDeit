// Dom7
var $$ = Dom7;

var poebota = [];
// Framework7 App main instance
var app = new Framework7({
    root: '#app', // App root element
    id: 'io.framework7.testapp', // App bundle ID
    name: 'Framework7', // App name
    theme: 'auto', // Automatic theme detection
    // App root data
    data: function() {
        return {
            user: {
                firstName: 'John',
                lastName: 'Doe',
            },
        };
    },
    // App root methods
    methods: {
        helloWorld: function() {
            app.dialog.alert('Hello World!');
        },
    },
    // App routes
    routes: routes,
});

// Init/Create main view
var mainView = app.views.create('.view-main', {
    url: '/'
});

app.on('test3', function() {
    // document.querySelector("")

    /*   var formData = app.form.fillFromData('#thirdScreen', {
           usuallyEat: "Bananas! 🍌"
       });
       var secondScreenFormData = app.form.convertToData('#secondScreen');
       console.log("secondScreenFormData", secondScreenFormData)*/
    var formData = app.form.convertToData('#thirdScreen');
    app.request.get('http://127.0.0.1:7777/Ings?h=' + formData.usuallyEat, function(data) {
        console.log(data);
    });
});