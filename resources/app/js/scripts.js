let index = {
    init: () => {
        // Initialize the Preloader
        preload.init()

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function () {

            // Show homepage
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
            // Add links to buttons
            buttons.parseSerial()
        });
    },
    editSerial: (softwarename, year) => {
        // Create our message
        let message = {"name": "editSerial"};
        if (typeof softwarename !== "undefined" && typeof year !== 'undefined') {
            message.payload = [softwarename, year]
        }

        astilectron.sendMessage(message, (message) => {
            // check for errors
            if (message.name === "error") {
                console.log(message.payload);
                return
            }
            console.log(message.payload.html_string)
            // document.getElementById('content').innerHTML = message.payload.html_string
        });
    }
};