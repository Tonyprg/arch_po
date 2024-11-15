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

async function click_event () {
    json = await get_json('http://localhost:8080/bingo');
    console.log(json);
}
