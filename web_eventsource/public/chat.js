window.addEventListener('load', () => {
    if (!window.EventSource) {
        alert("No EventSource!")
    }

    const chatLog = document.getElementById("chat-log")

    isBlank = (string) => {
        return string == null || string.trim() === "";
    };

    let username;
    while (isBlank(username)) {
        username = prompt("What's your name?");
        if (!isBlank(username)) {
            document.getElementById('user-name').innerHTML = `<b>${username}</b>`;
        }
    }

    document.getElementById("input-form").addEventListener("submit", async (e) => {
        e.preventDefault();
        const chatMsg = document.getElementById("chat-msg")
        console.log(chatMsg.value, username)
        const formData = new FormData(document.getElementById("input-form"));
        formData.append('name', username);
        axios({
            method: "post",
            url: "/messages",
            data: formData,
            headers: { "Content-Type": "multipart/form-data" },
        })
        .then(function (response) {
            //handle success
            console.log(response);
        })
        .catch(function (response) {
            //handle error
            console.log(response);
        });

        chatMsg.value = "";
        chatMsg.focus();
        return;
    });

    const addMessage = (data) => {
        let text = "";
        if (!isBlank(data.name)) {
            text = '<strong>' + data.name + ': </strong>';
        }
        text += data.msg
        document.getElementById('chat-log').insertAdjacentHTML('beforeend', '<div><span>' + text + '</span></div>');
    }

    const es = new EventSource('/stream');
    es.onopen = () => {
        axios({
            method: "post",
            url: "/users",
            data: {
                name: username
            },
            headers: { "Content-Type": "multipart/form-data" },
        })
        .then(function (response) {
            //handle success
            console.log(response);
        })
        .catch(function (response) {
            //handle error
            console.log(response);
        });
    }
    es.onmessage = function (e) {
        const msg = JSON.parse(e.data);
        addMessage(msg);
    }

    window.onbeforeunload = function () {
        axios({
            method: "delete",
            url: "/users?username=" + username,
        })
        .then(function (response) {
            //handle success
            console.log(response);
        })
        .catch(function (response) {
            //handle error
            console.log(response);
        });

    }
    
})
