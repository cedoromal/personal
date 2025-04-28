const conn = new WebSocket((document.location.protocol === "https:" ? "wss" : "ws" + "://" + document.location.host + "/live-reload"));

conn.onopen = () => {
    console.log("Connected")
}

conn.onclose = () => {
    console.log("Connection Closed")
    setTimeout(() => {
        fetch(document.location.origin + "/health").then((_) => {
            console.log("Reloading...")
            location.reload();
        })
    }, 1000);
};