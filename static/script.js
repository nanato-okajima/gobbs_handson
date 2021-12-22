let socket = null;

document.addEventListener('DOMContentLoaded', function(){
    socket = new WebSocket("ws://127.0.0.1:8080/ws")

    socket.onopen = () => {
        console.log("Successfully connectted!")
    }
})
