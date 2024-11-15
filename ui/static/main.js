async function get_json (url)
{
    let response = await fetch(url);
    if (response.ok)
    {
        let json = await response.json();
        return json;
    }
    else
    {
        alert("Ошибка HTTP: " + response.status);
    }
}

// function click_event ()
// {
//     fetch ('http://localhost:8080/bingo')
//         .then((response) => {
//             return response.json();
//         })
//         .then((data) => {
//             console.log(data);
//             console.log(data["name"])
//         });
// }

async function click_event () {
    json = await get_json('http://localhost:8080/bingo');
    console.log(json);
}
