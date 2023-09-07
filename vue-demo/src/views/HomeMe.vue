<template>
  <div>
  <el-form 
  :inline="true" 
  :model="user" 
  label-width="100px" 
  ref="uploadform"
  enctype="multipart/form-data"

  class="commit-form">
    <div id="CommitMes"><h1>{{this.tableData.Sname}}的作业</h1></div>

    
       <el-table :data="tableData" style="width: 100%">
          <el-table-column type="selection" width="55"></el-table-column>
          <el-table-column type="index" width="50"></el-table-column>


          <!--<el-table-column prop="Cno" label="课程号" align="center"></el-table-column>-->
          <el-table-column prop="Course" label="课程名" align="center"></el-table-column>

          <el-table-column prop="Sno" label="学号" align="center"></el-table-column>

          <el-table-column prop="Title" label="作业名" align="center"></el-table-column>
          
          <el-table-column prop="Deadline" label="截止日期" align="center"></el-table-column>

          <el-table-column prop="Status" align="center" label="提交状态">
            <div>
              {{Status ? '已提交' : '未提交'}}
            </div>
          </el-table-column>

          <el-table-column prop="operate" align="center" label="操作">
            <template slot-scope="scope">           
              <el-button  type="primary"  @click="submitAssignment(scope.row)">上传作业</el-button>
              <el-button  type="primary"  @click="submitVisible=true">修改作业</el-button>          
            </template>
              </el-table-column>
      </el-table>      
   </el-form>
   <el-dialog  v-if="submitVisible" title="上传作业" :visible.sync="submitVisible" width="50%">
              <!--<el-button type="primary" @click="submitAssignment()">-->
               
                
                  <!--<el-upload
                    class="upload-demo"
                    :limit="1"
                    :on-success="submitAssignment"
                    :before-upload="beforeUpload"
                    :action = 'uploadActionUrl'
                    :file-list="fileList">
                    <el-button type="primary">
                      提交作业
                    </el-button>
                  </el-upload>-->
                  <el-upload
                      :action='uploadActionUrl'
                      :before-upload="beforeUpload"
                      :limit="1"                   >  
                      
                      <el-button slot="trigger" type="primary"> 选择文件</el-button>
                      <!--<el-button slot="trigger" type="primary" @click="submitAssignment()"> 上传文件</el-button>-->
                    </el-upload>
                
              
  </el-dialog>

  </div>
</template>

<script>
import api from '@/api/API';
import axios from 'axios';

export default {
  name:'CommitMes',
  created() {
      const User = JSON.parse(localStorage.getItem('user'))
      console.log(" 哈哈哈  ",User)
      this.tableData = User
      this.user = User
      console.log(" 哈哈哈  ",this.tableData)
    },
  data() {
    return {
      submitVisible:false,
      token: 'your_token',
    fileList: [], // 上传的文件列表
    uploading: false, // 是否正在上传
    uploadActionUrl: 'http://172.29.88.83:9999/upload',
      user:{
          Course:'',
          Cno:'',
          Sno:'',
          Title:'',
          Deadline:'',
          Status:'',
      },
      FormData:new FormData(),
      tableData:[{
          Title:'实验一',
          Cno:'',
          Course:'数据库',
          Sno:'',
          deadline:'2023-6-1',
          Status:'已提交',
      }]
    };
  }, 
  methods: {  
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
        url:'http://172.29.88.83:9999/upload',
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

