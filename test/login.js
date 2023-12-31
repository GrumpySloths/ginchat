const login = () => {
    const formData = new FormData();
    formData.append("name", "阿库亚")
    formData.append("passwd", "123")
    formData.append("rePasswd", "123")
    const request = new Request("http://localhost:8080/user/CreateUser", {
        method: "POST",
        body: formData,
    });
    const url = request.url;
    console.log(request.body)
    const method = request.method;
    const credentials = request.credentials;
    const bodyUsed = request.bodyUsed;

    fetch(request)
        .then((response) => {
            if (response.status === 200) {
                return response.json();
            } else {
                throw new Error("Something went wrong on API server!");
            }
        })
        .then((response) => {
            console.debug(response);
            // …
        })
        .catch((error) => {
            console.error(error);
        });
}
login()
