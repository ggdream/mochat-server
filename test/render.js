let conn = null


function create(from, to) {
    const url = `ws://localhost:8080/chat/websocket?from=${from}&to=${to}`
    conn = new WebSocket(url)


    conn.onopen = () => {
        console.log('成功连接！')
    }

    conn.onerror = e => {
        console.error(e)
    }

    conn.onclose = e => {
        console.log(e.code)
    }

    conn.onmessage = e => {
        console.log(e.data)
    }
}

function send() {
    const input = document.getElementById('root').firstElementChild
    conn.send(input.value)
    input.value = null
}


window.onload = () => {
    const from = prompt('Type your name:')
    const to = prompt('Type her name:')
    create(from, to)
}
