<template>
  <div class="login-container">
    <el-form 
    :inline="true" 
    :model="loginForm" 
    status-icon
    label-width="100px" 
    ref="form"
    :rules = "rules"
    margin="0 auto"
    class="login-form">
      <div class="LoginMes">登陆页面</div>

      <el-form-item label="用户名" size ="30px" prop="username">
        <el-input v-model="loginForm.username" placeholder="请输入用户名"></el-input>
      </el-form-item>

      <el-form-item label="密码" size="50px" prop="password">
        <el-input type="password" v-model="loginForm.password" placeholder="请输入密码" autocomplete="off" maxlength="16"></el-input>
      </el-form-item>

      <el-form-item label="验证码" size="30px" prop="code">
        <el-input v-model="loginForm.code" placeholder="请输入验证码"></el-input>
        <img :src="codeSrc" id="codeid" class="captcha-img" @click="refreshCode" style="width: 100px; height: 40px;">
      </el-form-item>

      <el-form-item label="身份" prop="identity">
        <div style="text-align: center">
        <el-radio v-model="loginForm.identity"  label="0">管理员</el-radio>
        <el-radio v-model="loginForm.identity"  label="1">学生</el-radio>
        <el-radio v-model="loginForm.identity"  label="2">教师</el-radio>
        </div>
      </el-form-item>

      <el-form-item style="text-align: center">
        <el-button type="primary" @click = "onsubmit" >登录</el-button>
        <el-button type= "success" @click = "register">注册</el-button>
      </el-form-item>

      
      <!--<el-row style="text-align: center">-->
        <a href="/changepassword" title="忘记密码">忘记密码 </a>
        
        <el-dialog 
          title="修改密码" 
          :visible.sync="changeVisible" 
          width="30%">
            <el-form ref="form" :model="changeForm" label-width="80px">
              //绑定的模型和登录的一样的话，点击弹窗就会记录有用户名和旧密码
              <el-form-item label="学号">
                  <el-input v-model="changeForm.username"></el-input>
                </el-form-item>

                <el-form-item label="旧密码">
                  <el-input v-model="changeForm.oldpassword"></el-input>
                </el-form-item>
                
                <el-form-item label="新密码">
                  <el-input v-model="changeForm.newpassword"></el-input>
                </el-form-item>

                <el-form-item label="确认新密码" type="password">
                  <el-input v-model="changeForm.comfirmpassword"></el-input>
                </el-form-item>

                <el-form-item label="邮箱" type="password">
                  <el-input v-model="changeForm.email"></el-input>
                  <el-button type="primary" @click="SendEmail()">发送验证码</el-button>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="changePass">提交</el-button>
                    <el-button @click="changeVisible = false">取消</el-button>
                </el-form-item>

            </el-form>
        </el-dialog>
      

    </el-form>
  </div>
</template>

<script>

import axios from 'axios'
import globalurl from '../views/GlobalMe.vue'
  export default {
    name:'LoginMe',
    created(){
      this.refreshCode();
    },
    data() {
      return {
        changeVisible:false,
        codeSrc: '', // 验证码图片地址
        codeId:'',
        EmailCode:'',//邮箱验证码
        changeForm:{
          username:'',
          oldpassword:'',
          newpassword:'',
          comfirmpassword:'',
          email:'',
        },
        loginForm: {
          username: '2015200002',
          password: '059219',
          code:'',
          identity:'1',
          codeid:this.id,
        },
        FormData:new FormData(),
        
        rules: {
        username: [{required: true, message: "请输入账号", trigger: "blur"},
        {
            validator: (rule, value, callback) => {
              //const reg = /^[\u4E00-\u9FA5A-Za-z0-9]+$/  //正则表达式 只能输入英文、汉字与数字
              const reg = /^[0-9]+$/
              if (!reg.test(value)) {
                callback(new Error('请勿输入特殊字符'))
              } else {
                callback()
              }
            }
          }
        ]
        },
        password:[
          {required: true, message: "请输入密码", trigger: "blur"},
          { min: 0, max: 16, message: '长度在 0 到 16个字符', trigger: 'blur' }
        ],
        code:[{ required: true, message: '请输入验证码', trigger: 'blur' }]
      }      
    },
    mounted(){
      //this.refreshCode()  //获取验证码
    },
    methods: {
      
      // 刷新验证码图片
    refreshCode () {
      axios({
        method:'get',
       
        url:globalurl.BASEURL+'getImage',
      }).then(res=>{
       
        this.codeSrc =res.data.cpatchImg 
        this.codeId = res.data.captchaID
      })
    },
    submit(){
      if (this.loginForm.code === this.codeSrc){
        onsubmit
      }else(
        alert("验证码错误"),
        this.resetForm("form")
        //window.location.href="/login"
      )
    },
      onsubmit() {
          var _this = this
          _this.loginForm.codeid = this.codeId
          axios({
            method:'post',
            url:globalurl.BASEURL+'do_login',
            data:_this.loginForm,
            Headers:{"Content-Type":"multipart/FormData" },
          }).then(res => res.data).then(res => {
                
          if (res.succession === 'true') {
            if (_this.loginForm.identity==0){
              window.location.href = "/admin";

            }else if(_this.loginForm.identity==1) {
              console.log("email",res.email)
              if (res.email !== "null"){
                //localStorage.setItem('user', JSON.stringify(res.data))
                localStorage.setItem('username',_this.loginForm.username)
                localStorage.setItem('password',_this.loginForm.password)
                window.location.href = "/stu";
              }else{                
                window.location.href = "/bindemail";
              }              
           }else{
              if (res.email !== "null"){
                console.log(JSON.stringify(res.data))
                console.log("asac",res.email)
                localStorage.setItem('username',_this.loginForm.username)
                window.location.href = "/teacher";
              }else{               
                window.location.href = "/bindemail";
              }
            }
          } else if(res.code === 400){
            alert("该账号未注册")//注册失败，返回注册页面
          }else{
            this.refreshCode();
            alert("密码错误")
          }
        })
      },
      register(){
        this.$router.replace('/register')
      },
      SendEmail(){
        axios({
          method:'post',
          url:globalurl.BASEURL+'SendEmail',
          data:this.changeForm
        }).then(res=>{
          this.EmailCode=res.data.data
          console.log("code=",this.EmailCode)
        })
      },
      changePass(){
        this.changeVisible = true;
        if (this.changeForm.newpassword != this.changeForm.comfirmpassword){
          alert("前后密码不正确");
        }
      },
    }
  }

</script>

<style lang="scss" scoped>
.LoginMes {
  text-align: center;
  color: #f3730a;
  font-family: 幼圆;
  font-size: 60px;
  margin-bottom: 40px;
}
.login-container {
  /*position: absolute;*/
  width: 100%;
  height: 100%;
  overflow: hidden;
  background-image: url('../assets/images/R.jpg');
  background-size: 100% 100%;
  //background-repeat: no-repeat;
}


.login-form {
  width: 480px;
  margin: 215px auto;
  background-color: rgb(143, 180, 229);
  padding: 50px;
  border-radius: 80px;
}
  
</style>

