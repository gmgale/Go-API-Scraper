//  These tests are for testing on Postman
//  pm.environment.get("Id=threads") receives the enviroment varible
//  initally set to 0

let threads = parseInt(pm.environment.get("Id=threads"));

pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Content-Type is text/plain", function () {
    pm.response.to.have.header('Content-Type');
    pm.expect(pm.response.headers.get('Content-Type')).to.include('text/plain');
});

pm.test("Number of threads is displayed and correct", function () {
    if (isNaN(threads)){
        pm.expect(pm.response.text()).to.include("Invalid input (" + threads + ")");

    }else if (threads == 0) {
        pm.expect(pm.response.text()).to.include("Threads cannot be 0");
    }
    else{
        pm.expect(pm.response.text()).to.include("Threads" );
    }
});

pm.environment.set("Id=threads", threads+1);