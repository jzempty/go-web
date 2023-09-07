<template>
  <div class="page-register">
     

      <el-form
        ref="bindform"
        :model="ruleForm"
        :rules="rules"
        label-width="100px"
        class="demo-ruleForm"
        autocomplete="off"
        style="width: 400px; margin: 0 auto;"
        
      >
        
          <el-form-item label="用户名" prop="username" width="200px">
            <el-input v-model="ruleForm.username" width="30px"/>
          </el-form-item>
          <el-form-item label="旧密码" prop="oldpassword" minlength="6" maxlength="16">
            <el-input v-model="ruleForm.oldpassword" type="password" />
          </el-form-item>
          <el-form-item label="新密码" prop="newpassword" minlength="6" maxlength="16">
            <el-input v-model="ruleForm.newpassword" type="password" />
          </el-form-item>
          <el-form-item label="确认密码" prop="comfirmpassword" minlength="6" maxlength="16">
            <el-input v-model="ruleForm.comfirmpassword" type="password" />
          </el-form-item>
          <el-form-item label="身份" prop="identity">
            <div style="text-align: center">
            <el-radio v-model="ruleForm.identity"  label="0">管理员</el-radio>
            <el-radio v-model="ruleForm.identity"  label="1">学生</el-radio>
            <el-radio v-model="ruleForm.identity"  label="2">教师</el-radio>
            </div>
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="ruleForm.email" />
            <el-button :loading="codeLoading" :disabled="isDisable" size="mini" align="left" round @click="sendMsg">发送验证码</el-button>
            <span class="status">{{ statusMsg }}</span>
          </el-form-item>
          <el-form-item align="right" label="验证码" prop="code" max-width="100px">
            <el-input v-model="ruleForm.code" maxlength="6" />
          </el-form-item>
          <el-form-item >
            <el-button :loading="codeLoading" :disabled="isDisable" size="mini" round @click="submitForm()">提交</el-button>
          </el-form-item>
        
      </el-form>
  </div>
</template>

<script>
// import { getEmailCode } from '@/api'
import axios from 'axios';
import globalurl from './GlobalMe.vue'
//import { threadId } from 'worker_threads';
export default {
  data() {
    return {
      active: 0,
      statusMsg: '',
      error: '',
      isDisable: false,
      codeLoading: false,
      Ecode:'',
      ruleForm: {
        code: '',                 //邮件验证码
        username: "",
        oldpassword:'',
        newpassword: "",
        comfirmpassword: "",
        email: "",
        Status:'',
        identity:'',
      },
      rules: {
        email: [{
          required: true,
          type: 'email',
          message: '请输入邮箱',
          trigger: 'blur'
        }],
        code: [{
          required: true,
          message: '验证码',
          trigger: 'blur'
        }, {
          validator: (rule, value, callback) => {
            if (value === '') {
              callback(new Error('验证码不能为空'))
            } else if (value != this.Ecode) {
              callback(new Error('验证码不正确，请重新获取'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        }],
        newpassword: [{
          required: true,
          message: '创建密码',
          trigger: 'blur'
          }, { pattern: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,16}$/, message: '密码必须同时包含数字与字母,且长度为 6-16位' }],
        comfirmpassword: [
          {
          required: true,
          message: '确认密码',
          trigger: 'blur'
          }, 
          {
          validator: (rule, value, callback) => {
            if (value == '') {
              callback(new Error('请再次输入密码'))
            } else if (value != this.ruleForm.newpassword) {
              console.log(this.ruleForm.newpassword,value)
              callback(new Error('两次输入密码不一致'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        }]
      }
  }
  
  },
  
  methods: {
    sendMsg() {
      /*const self = this
      let namePass
      let emailPass
      let timerid
      console.log(timerid)
      if (timerid) {
        return false
      }
      this.$refs['ruleForm'].validateField('name', (valid) => {
        namePass = valid
      })
      self.statusMsg = ''
      if (namePass) {
        return false
      }
      this.$refs['ruleForm'].validateField('email', (valid) => {
        emailPass = valid
      })
      // 向后台API验证码发送
      if (!namePass && !emailPass) {
        self.codeLoading = true
        self.statusMsg = '验证码发送中...'
        getEmailCode(self.ruleForm.email).then(() => {
          this.$message({
            showClose: true,
            message: '发送成功，验证码有效期5分钟',
            type: 'success'
          })
          let count = 60
          self.ruleForm.code = ''
          self.codeLoading = false
          self.isDisable = true
          self.statusMsg = `验证码已发送,${count--}秒后重新发送`
          timerid = window.setInterval(function() {
            self.statusMsg = `验证码已发送,${count--}秒后重新发送`
            if (count <= 0) {
              console.log('clear' + timerid)
              window.clearInterval(timerid)
              self.isDisable = false
              self.statusMsg = ''
            }
          }, 1000)
        }).catch(err => {
          this.isDisable = false
          this.statusMsg = ''
          this.codeLoading = false
          console.log(err.response.data.message)
        })
      }*/
    axios({
      method:'post',
      url: globalurl.BASEURL + "SendEmail",
      data:this.ruleForm,
      Headers:{"Content-Type":"multipart/FormData" },
    }).then(res=>{
      this.Ecode = res.data.code
      console.log("code",this.Ecode,res.data.code)
    })
    },

    submitForm(){
      if(this.ruleForm.username ===null || this.ruleForm.identity ===null){
        alert("请输入完整信息")
      }else if(this.ruleForm.code==''){
        alert("请输入验证码");
        return;
      }else if(this.Ecode != this.ruleForm.code){
        alert("验证码有误，请重新输入")
        this.$refs.bindform.resetFields();
        return;
      }else if(this.ruleForm.newpassword != this.ruleForm.comfirmpassword ||this.ruleForm.newpassword == ''){
        alert("两次密码输入不一样，请重新输入")
        this.$refs.bindform.resetFields()
      }else{
        this.ruleForm.Status="isok"
        axios({
        method:'post',
        url: globalurl.BASEURL + "BindEmail",
        data:this.ruleForm
      }).then(res=>{
        if (res.data.succession =='true'){
          alert("修改成功")
          window.location.href="/login"
        }else{
          alert("500")
          this.$refs.bindform.resetFields();
        }
      })
      }
      
    },

    
  }
};

</script>

<style  rel="stylesheet/scss" lang="scss">
.page-register {
  .header {
    //border-bottom: 20px solid rgb(235, 232, 232);
    min-width: 200px;
    color: #666;

    header {
      margin: 0 auto;
      padding: 10px 0;
      width: 980px;

      .login {
        float: right;
      }

      .bold {
        font-style: normal;
      }
    }
  }

  .register {
    color: #1890ff;
  }

  a {
    color: #1890ff;
    text-decoration: none;
    background-color: transparent;
    outline: none;
    cursor: pointer;
    transition: color 0.3s;
  }

  > section {
    margin: 0 auto 30px;
    padding-top: 30px;
    width: 20px;
    min-height: 30px;
    padding-right: 100px;
    box-sizing: border-box;

    .status {
      font-size: 12px;
      margin-left: 50px;
      color: #e6a23c;
    }

    .error {
      color: red;
    }
  }

  .footer {
    text-align: center;
  }
}
</style>

