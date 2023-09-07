<template>
    <el-form 
    :inline="true" 
    :model="user" 
    label-width="100px" 
    ref="form"
    margin="0 auto"
    class="commit-form">
      <div id="CommitMes"><h1>{{User.Sname}}的作业</h1></div>

      
         <el-table
                :data="tableData"
                style="width: 100%"
                :row-class-name="tableRowClassName"
                @selection-change="handleSelectionChange" stripe>
            <el-table-column
                    type="selection"
                    width="55">
            </el-table-column>
            <el-table-column
                    type="index"
                    width="50">
            </el-table-column>


            <el-table-column prop="Course" label="课程名" align="center"></el-table-column>

            <el-table-column prop="Sno" label="学号" align="center"></el-table-column>

            <el-table-column prop="Title" label="作业名" align="center"></el-table-column>
            
            <el-table-column prop="Deadline" label="截止日期" align="center"></el-table-column>

            <el-table-column prop="Status" align="center" label="提交状态">
              <!--<div>
                {{Status ? '已提交' : '未提交'}}
              </div>-->
            </el-table-column>

            <el-table-column prop="operate" align="center" label="操作">
                <template #default="{ row }">
                <el-button v-if="!row.submitted" type="primary" @click="submitAssignment(row.workno)"><div>
                  <div>
                    <el-upload
                      class="upload-demo"
                      :limit="1"
                      :on-success="submitAssignment"
                      :before-upload="beforeUpload"
                      :action = 'uploadActionUrl'
                      :file-list="fileList">
                      <el-button type="primary">
                        提交作业
                      </el-button>
                    </el-upload>
                  </div>
                      </div><i class="el-icon-upload el-icon--right"></i></el-button>
                
                <el-button  type="primary" icon="el-icon-edit" @click="submitAssignment()">修改作业</el-button>
                </template>
            </el-table-column>
        </el-table>
        
     </el-form>

    
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
        token: 'your_token',
      fileList: [], // 上传的文件列表
      uploading: false, // 是否正在上传
      uploadActionUrl: 'http://172.29.194.191:9999/upload',
        user:{
            Course:'',
            Sno:'',
            Title:'',
            Deadline:'',
            Status:'',
        },
        FormData:new FormData(),
        tableData:[{
            Title:'实验一',
            Course:'数据库',
            Sno:'',
            deadline:'2023-6-1',
            Status:'已提交',
        }]
      };
    }, 
    methods: {  
    handleSuccess(response, file, fileList) {
      this.fileList = fileList
      alert("上传成功")
    },
      beforeUpload(file) {
        this.FormData.append('Course',this.user.Course)
        this.FormData.append('Sno',this.user.Sno)
        this.FormData.append('Title',this.user.Title)
        this.FormData.append('Status',this.user.Status)
        this.FormData.append('Deadline',this.form.Deadline)
        this.FormData.append('file',file)

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
    }
  },
      fetchAssignments() {
        api.getAssignments().then(assignments => {
          this.assignments = assignments;
        }).catch(error => {
          console.error(error);
        });
      },
      submitAssignment() {
        axios({
          method:'post',
          url:'http://172.29.194.191:9999/upload',
          data:this.commitForm
        }).then(),


        api.submitAssignment().then(() => {
          this.fetchAssignments();
          this.$message.success('作业提交成功！');
        }).catch(error => {
          console.error(error);
          this.$message.error('作业提交失败。请稍后重新尝试！');
        });
      },
    
  };
  </script>

