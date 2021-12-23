let socket = null;

document.addEventListener('DOMContentLoaded', function(){
    socket = new ReconnectingWebSocket(
        "ws://127.0.0.1:8080/ws",
        null,
        {
            debug: true,
            reconnectInterval: 3000
        })

    socket.onopen = () => {
        console.log("Successfully connectted!")
    }

    socket.onmessage = msg => {
        let postData = JSON.parse(msg.data)
        let postContainer = document.getElementById("post-container")

        switch (postData.action) {
            case "broadcast":
                let post = postData.post
                postContainer.innerHTML = postContainer.innerHTML + post
                break
        }
    }

    let userInput = document.getElementById("username")
    userInput.addEventListener("change", function() {
        let jsonData = {}
        jsonData["action"] = "username"
        jsonData["username"] = this.value
        socket.send(JSON.stringify(jsonData))
    })

    window.onbeforeunload = function() {
        console.log("User Leaving")
        let jsonData = {}
        jsonData["action"] = "left"
        socket.send(JSON.stringify(jsonData))
    }

    document.getElementById("post").addEventListener("keydown", function(e) {
        if (e.composed === "Enter") {
            if (!socket) {
                console.log("no connection")
                return
            }

            e.preventDefault()
            e.stopPropagation()
            sendPost()
        }
    })
})

function sendPost() {
    console.log("Send Post...")

    let jsonData = {}
    jsonData["action"] = "broadcast"
    jsonData["username"] = document.getElementById("username").value
    jsonData["post"] = document.getElementById("post").value

    socket.send(JSON.stringify(jsonData))

    document.getElementById("post").value = ""
}
