$(document).ready(function () {
    $("#loginform").submit(login);
});

function login(event) {
    let email = $("#email").val();
    let password = $("#password").val();
    var data = {
        "email":email,
        "password":password
    };
    loginUser(data).then(response => console.log(response))
    event.preventDefault();
}

loginUser = async (data) => {
    const settings = {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
        origin: "http://127.0.0.1:5500"
    };
    const fetchResponse = await fetch(`http://localhost:9091/login`, settings);
    const resdata = await fetchResponse.json();
    return resdata;
}



// let crud = [];

// const sendHttpRequest = (method, url, data) => {
//     const promise = new Promise((resolve, reject) => {
//         const xhr = new XMLHttpRequest();
//         xhr.open(method, url);

//         xhr.responseType = 'json';

//         if (data) {
//             xhr.setRequestHeader('Content-Type', 'application/json');
//             xhr.setRequestHeader('Access-Control-Allow-Origin', '*');
//         }

//         xhr.onload = () => {
//             if (xhr.status >= 400) {
//                 reject(xhr.response);
//             } else {
//                 resolve(xhr.response);
//             }
//         };

//         xhr.send(JSON.stringify(data));
//     });
//     return promise;
// };

// function manageData() {

//     let fname = $("#fname").val();
//     let sname = $("#sname").val();
//     let email = $("#email").val();
//     let dob = $('#dob').val();
//     let password = $("#password").val();

//     crud.push({ fname, sname, email, dob, password });
//     setCrudData(crud)

//     $(".form-control").val("");
// }

// function populateData(arr) {
//     if (arr != null) {
//         let html = '';
//         for (let i = 0; i < arr.length; i++) {
//             html = html + `<tr><td>${i + 1}</td><td>${arr[i].id}</td><td>${arr[i].fname}</td><td>${arr[i].lname}</td><td>${arr[i].email}</td>
//             <td>${arr[i].dob}</td><td><a href="javascript:void(0)" onclick="editData(${i})">Edit
//             </a>&nbsp;&nbsp;<a href="javascript:void(0)" onclick="deleteData(${i})">Delete</a></td></tr>`;
//         }
//         $("#root").html(html);
//     }
//     $('#myTable').DataTable();
// }

// function deleteData(rid) {
//     let arr = getCrudData();
//     arr.splice(rid, 1);
//     crud.splice(rid, 1);
//     setCrudData(arr);
//     populateData();
//     location.reload();
// }

// function getCrudData() {
//     sendHttpRequest('GET', 'http://localhost:9091/getData').then(responseData => {
//         populateData(responseData);
//     }).catch(err => {
//         console.log(err);
//     });
// }

// function setCrudData(arr) {
//     sendHttpRequest('POST', 'http://localhost:9091/setData', arr).then(responseData => {
//         console.log(responseData);
//     }).catch(err => {
//         console.log(err);
//     });
// }

// function editData(rid) {
//     id = rid;
//     let arr = getCrudData();
//     $("#employeeName").val(arr[rid].name);
//     $("#employeeCode").val(arr[rid].code);
//     $("#employeeEmail").val(arr[rid].email);
//     $("#employeeDesignation").val(arr[rid].desig);
//     $("#employeeDOB").val(arr[rid].dob);
//     $("#update").html("Update");
//     $('#myModal').modal('show');
// }
