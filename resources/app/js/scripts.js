console.log("Script Loaded.")
let index = {
    init: () => {
        // Initialize the Preloader
        preload.init()

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function () {
            // Listen
            index.listen();

            // Test homepage
            index.software();
        })
    },
    software: () => {
        // Create our message
        let message = {"name": "software"};

        preload.show()
        // Send our message
        astilectron.sendMessage(message, (message) => {

            // check for errors
            if (message.name === "error") {
                console.log(message.payload);
                return
            }
            preload.hide()
            document.getElementById('content').innerHTML = message.payload.html_string
        });
    },
    listen: () => {
        astilectron.onMessage(function (message) {
            switch (message.name) {
                case "about":
                    return {payload: "payload"};
                case "check.out.menu":
                    console.log(message.payload);
                    break;
            }
        });
    },
    editSerial: () => {
        // Create our message
        let message = {"name": "editSerial"};

        astilectron.sendMessage(message, (message) => {

            // check for errors
            if (message.name === "error") {
                console.log(message.payload);
                return
            }
            document.getElementById('content').innerHTML = message.payload.html_string
        });
    }
};