<template>
  <div>
    <div id="CommitMes"><h1>欢迎您：{{this.username}}</h1></div>
      <el-button align="rigth" type="primary" class="el-icon-back" plain @click="LoginOut()">安全退出</el-button>

    <!--搜索表单-->
    <el-form :inline="true" :model="searchCourse" class="demo-form-inline">
            <el-form-item label="课程名">
                <el-input v-model="searchCourse.Course" placeholder="课程名称"></el-input>
            </el-form-item>   
                    
            <el-form-item label="提交状态">
                <el-radio-group v-model="searchCourse.Status">
                    <el-radio label="已提交"></el-radio>
                    <el-radio label="未提交"></el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" class="el-icon-view" @click="Search()">查询</el-button>
            </el-form-item>
        </el-form>
        <el-row align="left" >
          <el-button type="primary" size="medium" @click="selectAll()">查询所有</el-button>
        </el-row>
  <el-form 
  :inline="true" 
  :model="user" 
  label-width="100px" 
  ref="uploadform"
  enctype="multipart/form-data"
  class="commit-form">
       <el-table ref="SZJ" :data="tableData" style="width: 100%">
          <el-table-column type="selection" width="55"></el-table-column>
          <el-table-column type="index" width="50"></el-table-column>


          <!--<el-table-column prop="Cno" label="课程号" align="center"></el-table-column>-->
          <el-table-column prop="Course" label="课程名" align="center"></el-table-column>

          <el-table-column prop="Sno" label="学号" align="center"></el-table-column>

          <el-table-column prop="Aid" label="作业号" align="center"></el-table-column>
          <el-table-column prop="Title" label="作业名" align="center"></el-table-column>
          
          <el-table-column prop="Deadline" label="截止日期" align="center"></el-table-column>

          <el-table-column prop="Status" align="center" label="提交状态">
            <!--<div>
              {{Status ? '已提交' : '未提交' }}
            </div>-->
          </el-table-column>

          <el-table-column prop="Score" label="成绩" align="center"></el-table-column>

          <el-table-column prop="operate" align="center" label="操作">
            <template slot-scope="{row}">          
              <el-button  type="primary"  @click="Openform(row)">上传作业</el-button>
              <el-button  type="primary" class="el-icon-view"  @click="submitVisible=true">预览作业</el-button>          
            </template>
              </el-table-column>
      </el-table>      
   </el-form>
   
   <el-dialog title="上传作业" :visible.sync="submitVisible" width="50%">
    <el-form ref="uploadForm" :model="user" :rules="importFormRules" label-width="130px">
        <el-form-item label="课程名" prop="Course" >
          <el-input v-model="user.Course" ></el-input>
        </el-form-item>
        <el-form-item label="学号" prop="Sno" >
          <el-input v-model="user.Sno"></el-input>
        </el-form-item>
        <el-form-item label="作业名" prop="Title" >
          <el-input v-model="user.Title"></el-input>
        </el-form-item>
        <el-form-item label="上传文件:" prop="file">
          <el-upload
            class="upload-demo"
            ref="upload"
            :before-upload="beforeUpload"
            :http-request="httpRequest"
            :limit="1">
            <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
            <div slot="tip" class="el-upload__tip">只能上传pdf文件，且不超过8M</div>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitImportForm">开始导入</el-button>
          <el-button type="info" @click="submitVisible = false">关闭窗口</el-button>
        </el-form-item>
      </el-form>
              
  </el-dialog>

  </div>
</template>

<script>
import api from '@/api/API';
import axios from 'axios';
import globalurl from '../views/GlobalMe.vue'

export default {
  name:'CommitMes',
  created() {
      this.firstselectAll();
      this.selectAll();
    },
  data() {
    return {
      submitVisible:false,
      token: 'your_token',
      username:'',
      password:'',
      Aid:'',
    fileList: [], // 上传的文件列表
    uploading: false, // 是否正在上传
    uploadActionUrl: globalurl.BASEURL+'upload',
    searchCourse:{
      Course:'',
      Status:'',
    },
    nowuser:[],
      
      user:{
          Course:'',
          Cno:'',
          Sno:'',
          Title:'',
          Deadline:'',
          Status:'',
          username:'',
          Score:'',
          Aid:'',
      },
      login:{},
      FormData:new FormData(),
      tableData:[{
          Title:'实验一',
          Cno:'',
          Course:'数据库',
          Sno:'',
          Deadline:'2023-6-1',
          Status:'已提交',
          username:'',
          Score:'',
          Aid:'',
      }]
    };
  }, 
  methods: { 
    LoginOut(){
      window.sessionStorage.clear()
      this.$router.push('/login')
    }, 
    firstselectAll(){
    
      this.username  = localStorage.getItem('username')
      this.password = localStorage.getItem('password')
      
    },
    selectAll(){
      const login = {
        username:this.username,
        password:this.password
      }
      this.login = login
      console.log("sa",login)
      axios({
        method:'post',
        url:globalurl.BASEURL+'StudentAssignmentQuery',
        data:this.login
      }).then(res=> {
         this.tableData = res.data.data;
         this.user = res.data.data
         this.Aid = this.user.Aid
         console.log("dayin",this.user,res.data.data,this.tableData)
         for (let i = 0;i<this.tableData.length; i++) {
        if(this.tableData[i].Status===true){
        this.tableData[i].Status='已提交'
        if(this.tableData[i].Score==-1){
        this.tableData[i].Score='尚未批改'
        }
        }else{
        this.tableData[i].Status='未提交'
        this.tableData[i].Score='请先提交作业'
        }
        }
         })
    },
    Search(){
     this.nowuser=[]
     for(let i=0;i<this.user.length;i++){
      if(this.searchCourse.Course.length>0){
        if (this.user[i].Status == this.searchCourse.Status && this.user[i].Course==this.searchCourse.Course){
        this.nowuser.push(this.user[i])
        }
      }else{
        if (this.user[i].Status == this.searchCourse.Status){
        this.nowuser.push(this.user[i])
        }
      }
      
     }
     this.tableData = this.nowuser
    },
    Openform(row){
      this.user.Course = row.Course
      this.user.Sno=row.Sno
      this.user.Title = row.Title
      this.user.Aid = row.Aid
      this.submitVisible=true
    },
    handleFileChange(file,fileList){
      this.fileList = fileList
    },
    
    /*submitAssignment() {
      // 构建 FormData 数据
      const formData = new FormData();
      for (const file of this.fileList) {
        formData.append('file', file);
      }
      formData.append('table_data', JSON.stringify(this.tableData)); // 注意这里需要把表格数据转换成 JSON 字符串

      // 发送表单请求
      axios({
        method:'post',
        url:'http://172.29.88.83:9999/upload',
        body: formData,
        headers:{'Content-Type': 'multipart/form-data', },
      }).then( res=> {
          if (res.succession === true){
            alert("sss")
          }
        })
        /*.catch(error => {
          // 异常处理逻辑
        });
    },*/
    handleSuccess(response, file, fileList) {
      this.fileList = fileList
      alert("上传成功")
    },
  // 覆盖默认的上传行为，可以自定义上传的实现，将上传的文件依次添加到fileList数组中,支持多个文件
    httpRequest(option) {
          this.fileList.push(option)
        },
    beforeUpload(file) {
      /*this.FormData.append('Course',this.user.Course)
      this.FormData.append('Sno',this.user.Sno)
      this.FormData.append('Title',this.user.Title)
      this.FormData.append('Status',this.user.Status)
      this.FormData.append('Deadline',this.form.Deadline)
      this.FormData.append('file',file)*/

    const isPDF = file.type === 'application/pdf'
    const isLt10M = file.size / 1024 / 1024 < 10

    if (!isPDF) {
      this.$message.error('上传图片只能是 PDF 格式!')
    }
    if (!isLt10M) {
      this.$message.error('上传图片大小不能超过 8MB!')
    }
    return isPDF && isLt10M
    },
    submitImportForm() {
          // 使用form表单的数据格式
          const params = new FormData()
          // 将上传文件数组依次添加到参数paramsData中
          this.fileList.forEach((x) => {
            params.append('file', x.file)
          });
          // 将输入表单数据添加到params表单中
          //params.append('Course', this.user.Course)
          params.append('Sno', this.user.Sno)
          //params.append('Title', this.user.Title)
          params.append('Aid',this.user.Aid)
          console.log("AID",this.user.Aid)
          console.log("sad asca",params)
          //params.append('targetPassword', this.importForm.targetPassword)
      
      //这里根据自己封装的axios来进行调用后端接口
          axios({
            method:'post',
            //url:'http://localhost:9999/upload',
            url:globalurl.BASEURL+'upload',
            data:params,
            Headers:{"Content-Type":"multipart/FormData" },
          }).then(res=>res.data).then(res=>{
            if (res.succession ==='true'){
              this.selectAll();
  
              alert("上传成功")
              //this.tableData.Status = true;
              this.submitVisible=false;
            }else{
              alert("上传失败")
              this.submitVisible=false;
            }
          })
            
        },
    onSuccess(response, file, fileList) {
      console.log(response, file, fileList)
      this.$message.success('上传成功!')
    },
    submitAssignment1(tableData) {
        this.submitVisible=true
        this.submitAssignment1(tableData)
      }
      ,
    submitAssignment(User) {
        this.submitVisible=true
        axios({
          method:'post',
          url:globalurl.BASEURL+'upload',
          headers:{'Content-Type': 'application/json;charset=UTF-8'},
          data:User
        }).then(res=>res.data).then(res=>{
          console.log("宋湛杰",res)
        })
      },
    fetchAssignments() {
        api.getAssignments().then(assignments => {
          this.assignments = assignments;
        }).catch(error => {
          console.error(error);
        });
      }, 
  },
  
};
</script>

