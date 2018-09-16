// Dom7
var $$ = Dom7;

var poebota;
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
    on: {
        pageInit: function(page) {
            if (page.name === "result") {
                const template = page.$el.children("#tabletemplate").html();
                const compiled = Template7.compile(template);
                const content = compiled(poebota);
                console.log(content)
                page.$el.find("#resulttable").html(content);

                const template2 = page.$el.children("#tabletemplate2").html();
                const compiled2 = Template7.compile(template2);
                const content2 = compiled2(poebota);
                console.log(content2)
                page.$el.find("#resulttable2").html(content2);

                const template3 = page.$el.children("#tabletemplate3").html();
                const compiled3 = Template7.compile(template3);
                const content3 = compiled3(poebota);
                console.log(content3)
                page.$el.find("#resulttable3").html(content3);
            }
        }
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

app.on('test4', function() {
    window.open(poebota.LinkWithShops);
})

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
        poebota = JSON.parse(data);

        app.views.main.router.navigate("/result/")
    });
})