input = document.getElementById("Url")

async function callapi(event) {
    event.preventDefault()

    await fetch("http://localhost:8000/a", {
        mode: "no-cors",
        headers: {
            "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify ({
            value: input.value
        })
    })
    .then(response => response.json())
    .then(data => {
        // console.log(data)
        if (data.LongURL == "Wrong form") {
            document.getElementById("answer").innerHTML = "Your link was wrong!!"    
        }
        else {
            document.getElementById("answer").innerHTML = "http://localhost:8000/s/" + data.ShortURL
        }
    })

    input.value = ""
}   