var i = 1;

const sendHttpRequest = (method, url, data) => {
  return fetch(url, {
    method: method,
    body: JSON.stringify(data),
    headers: data ? { 'Content-Type': 'application/json' } : {}
  }).then(response => {
    return response.json();
  });
};

const get = () => {
  url = 'http://localhost:9091/users?page';
  res = url.concat("=",i)
  sendHttpRequest('GET',res).then(response =>{
      populateData(response.Data,response.Currpage,response.Lastpage)
  }).catch(err=>{
      window.alert(err)
  });
  i++;
}

const register = () => {
  fname = $('#fname').val();
  lname = $('#lname').val();
  email = $('#email').val();
  dob = $('#dob').val();
  password = $('#password').val();

  data = {
    fname:fname,
    lname:lname,
    email:email,
    dob:dob,
    password:password
  }

  sendHttpRequest('POST','http://localhost:9091/user',data).then(response =>{
    get();
  }).catch(
    window.alert('enter valid data')
  );
}

const dec = () => {
  i = i-2;
  get();
}

function populateData(arr,x,y) {
  if (arr != null) {
      let html = '';
      for (let i = 0; i < arr.length; i++) {
          html = html + `<tr><td>${arr[i].ID}</td><td>${arr[i].fname}</td><td>${arr[i].lname}</td><td>${arr[i].email}</td>
          <td>${arr[i].dob}</td><td><a href="javascript:void(0)" onclick="editData(${arr[i].ID})">Edit
          </a>&nbsp;&nbsp;<a href="javascript:void(0)" onclick="deleteData(${arr[i].ID})">Delete</a></td></tr>`;
      }
      $("#root").html(html);
  }
  info = "Showing page "+x+" of "+y
  $('#info').html(info)
}

// const login = (e) => {
//     e.preventDefault();
//     email = $('#email').val();
//     password = $('#password').val();
//     data = {
//         email:email,
//         password:password
//     }
//     sendHttpRequest('POST','http://localhost:9091/login',data).then(response =>{
//         window.alert(response.message)
//     }).catch(err=>{
//         window.alert(err.message)
//     });
// }

const deleteData = (id) => {
  url = 'http://localhost:9091/user';
  res = url.concat("/",id)
  sendHttpRequest('DELETE',res).then(response =>{
      get();
  }).catch(err=>{
      window.alert(err)
  });
}

const editData = (id) => {
  url = 'http://localhost:9091/user';
  res = url.concat("/",id)
  sendHttpRequest('PATCH',res).then(response =>{
      get();
  }).catch(err=>{
      window.alert(err)
  });
}

const search = () => {
  input = $('#myInput').val();
  url = 'http://localhost:9091/users?search';
  res = url.concat("=",input)
  sendHttpRequest('GET',res).then(response =>{
      populateData(response.Data,response.Currpage,response.Lastpage)
  }).catch(err=>{
      window.alert(err)
  });
}


getBtn = $('#get').click(get)
nextBtn = $('#next').click(get)
prevBtn = $('#prev').click(dec)
searchBtn = $('#search').click(search)
submitBtn = $('#submit').click(register)





