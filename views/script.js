$( document ).ready(function() {
  $('#updiv').hide();
});
var i = 1;
var Id = 1;

const sendHttpRequest = (method, url, data) => {
  return fetch(url, {
    method: method,
    body: JSON.stringify(data),
    headers: data ? { 'Content-Type': 'application/json' } : {}
  }).then(response => {
    if(response.ok){
      return response.json();
    }else{
      window.alert(response.statusText)
      return
    }
  }).catch(err =>{
    return err
  });
};

const get = () => {
  tkn = localStorage.getItem('token');
  url = 'http://localhost:9091/users?page';
  res = url.concat("=",i)
  res1 = res.concat("&token=",tkn)
  console.log(res1);
  sendHttpRequest('GET',res1).then(response =>{
      populateData(response.Data,response.Currpage,response.Lastpage)
  }).catch(err=>{
      console.log(err);
  });
  i++;
}

const register = (e) => {
  e.preventDefault();
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
  }).catch(err => {
    window.alert(err)
  })
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

const login = (e) => {
    e.preventDefault();
    email = $('#useremail').val();
    password = $('#userpass').val();
    data = {
        email:email,
        password:password
    }
    sendHttpRequest('POST','http://localhost:9091/login',data).then(response =>{
        localStorage.setItem('token',response)
        location.replace("http://127.0.0.1:5500/views/users.html")
    }).catch(err=>{
        window.alert(err)
    });
}

const deleteData = (id) => {
  url = 'http://localhost:9091/user';
  res = url.concat("/",id)
  sendHttpRequest('DELETE',res).then(response =>{
      window.alert(response)
      get();
  }).catch(err=>{
      window.alert(err)
  });
}

const editData = (id) => {
  $('#updiv').show();
  Id = id
}

const update = (e) => {
  e.preventDefault();
  Fname = $('#Fname').val();
  Lname = $('#Lname').val();
  Email = $('#Email').val();
  Dob = $('#Dob').val();

  data = {
    fname:Fname,
    lname:Lname,
    email:Email,
    dob:Dob
  }
  url = 'http://localhost:9091/user';
  res = url.concat("/",Id)
  sendHttpRequest('PATCH',res,data).then(response =>{
      $('#updiv').hide();
      window.alert(response)
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

const logout = () => {
  localStorage.removeItem("token");
  location.replace("http://127.0.0.1:5500/views/index.html")
  return
}


nextBtn = $('#next').click(get)
prevBtn = $('#prev').click(dec)
searchBtn = $('#search').click(search)
submitBtn = $('#submit').click(register)
updateBtn = $('#update').click(update)
getBtn = $('#getusers').click(get)
loginBtn = $('#loginsubmit').click(login)
logoutBtn = $('#logout').click(logout)






