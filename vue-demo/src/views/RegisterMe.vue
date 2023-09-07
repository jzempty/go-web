<template>
    <div class="register-container">
      <!-- 根标签 -->
      <el-form
          :model="registerform"
          status-icon
          :rules="rules"
          ref="registerform"
          label-width="100px"
          class="register-form">
   
        <h1 class="title">注 册</h1>

        <el-form-item label="姓名" prop="username">
          <el-input v-model="registerform.username" maxlength="20"></el-input>
        </el-form-item>

        <el-form-item label="学号" prop="useraccount">
          <el-input v-model="registerform.userno"
                    placeholder="账号作为登陆的唯一方式，一旦注册成功不可更改！"
                    maxlength="20"></el-input>
        </el-form-item>
   
        <el-form-item label="密码" prop="userpsd">
          <el-input
              type="password"
              v-model="registerform.password"
              placeholder="请输入密码"
              autocomplete="off"
              maxlength="16"
          ></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="checkPass">
          <el-input
              type="password"
              v-model="registerform.checkPassword"
              placeholder="请再次输入密码"
              autocomplete="off"
              maxlength="16"
          ></el-input>
        </el-form-item>

        <el-cascader  :options="registerform.options" ></el-cascader>
   
        <el-form-item style="text-align: center">
          <el-button type="primary" @click="submit('registerform')">提交</el-button>
        </el-form-item>
        <div class="reg-content">
            <h1>欢迎注册</h1>
            <span>已有帐号？</span> <a href="/login">登录</a>
        </div>
      </el-form>
    </div>
   
  </template>

<script>
import axios from 'axios'
export default {
  name: "RegisterMe",
  data() {
    return {
      registerform: {
        username: '1',
        userno: '2',
        password: '3',
        checkPassword:'4',
        options: [{
          value: 'shuyuan',
          label: '数学学院',
          children: [{
            value: '数学',
            label: '数学',
            children: [{
              value: 'shuxue',
              label: '数学201'
            }, {
              value: 'shuxue',
              label: '数学202'
            }]
          }, {
            value: 'xj',
            label: '信计',
            children: [{
              value: 'xji',
              label: '信计201'
            },  {
              value: 'xji',
              label: '信计202'
            }]
          }, {
            value: 'xinan',
            label: '信安',
            children: [{
              value: 'xa',
              label: '信安201'
            }, {
              value: 'xa',
              label: '信安202'
            }]
          }]
        }, {
          value: 'qita',
          label: '其他学院',
          children: [{
            value: 'huanyuan',
            label: '环院',
            children:[{
                value:"huanke",
                label:"环科"
            },{
                value:"huangong",
                label:"环工"
            }]
          }, {
            value: 'tiyuan',
            label: '体院'
          }, {
            value: 'tumu',
            label: '土木'
          }]
        }],
    },
        defaultProps: {
          children: 'children',
          label: 'label'
        },
     
      rules: {
        userno: [{required: true, message: "请输入账号", trigger: "blur"},
          {min: 1, max: 20, message: '请勿超过20个字符！', trigger: 'blur'},
          {validator: (rule, value, callback) => {
              const reg = /^[0-9]+$/  //正则表达式 只能输入英文、汉字与数字
              if (!reg.test(value)) {
                callback(new Error('请勿输入特殊字符'))
              } else {
                callback()
              }}}
        ],
        password: [{required: true, message: "请输入密码", trigger: "blur"},
          {min: 0, max: 16, message: '长度在 0 到 16个字符', trigger: 'blur'}
        ],
        checkPassword: [{required: true, message: "请再次输入密码", trigger: "blur"},
          {min: 0, max: 16, message: '长度在 0 到 16个字符', trigger: 'blur'},
          {validator: (rule, value, callback) => {
              if (value !== this.registerform.password) {
                callback(new Error("两次输入密码不一致"));
              } else {
                callback()
              }}}]
      },
    }
  },
  methods: {
    handleNodeClick(data) {
        console.log(data);
      },
    submit() {
      //校验输入
      this.$refs.registerform.validate((valid)=>{
        if (valid){
         //注册功能
          axios.post('http://127.0.0.1:9999/do_register', this.registerform).then(res => res.data).then(res => {
          console.log(res)
        //用户注册成功或设计师成功提交材料后，显示相关信息后自动返回首页
        if (res.code === 200) {
          if (this.form.roleId==1){
            alert("用户注册成功")
            window.location.href = "/login";
          }else{
            alert("您的注册材料已提交，请等候管理员审批！")
            window.location.href = "/login";
          }
        } else {
          alert("该账号已注册过！请重新输入")//注册失败，返回注册页面
        }
      })
        }else{
          alert("请按照规范输入账号")
        }
      })
    }    
}
}
 
</script>

<style scoped>
.register-container {
  /*position: absolute;*/
  width: 100%;
  height: 100%;
  overflow: hidden;
 background-image:none;
  background-size: 100% 100%;
  background-repeat: no-repeat;
}
.register-form {
  width: 600px;
  margin: 150px auto;
  background-color: rgb(214, 174, 218);
  padding: 30px;
  border-radius: 20px;
}
.title {
  text-align: center;
  font-family: 幼圆;
}
 
 
</style>