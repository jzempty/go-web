<!-- eslint-disable no-dupe-keys -->
<template>
  <div id="TeacherPage" class="container">
    <div><h2>管理实验课程</h2></div>        

        <el-row >
          <el-button align="rigth" type="primary" class="el-icon-back" plain @click="LoginOut()">安全退出</el-button>          
        </el-row>

        <el-row>
          <el-button type="primary" plain @click="addcourse=true">创建课程</el-button>
          <el-button type="text" @click="selectAll()">查询所有</el-button>
          
        </el-row>
        <el-dialog title="添加课程" :visible.sync="addcourse" width="50%">
          <el-form ref="form" :model="MgForm" label-width="120px" class="form-container">
            <el-form-item label="课程名称">
              <el-input v-model="MgForm.CourseName" placeholder="请输入课程名称"></el-input>
            </el-form-item>
            <el-form-item label="开始时间">
              <el-date-picker v-model="MgForm.StartDate" type="datetime" placeholder="选择日期时间"></el-date-picker>
            </el-form-item>
            <el-form-item label="课程学分">
              <el-input v-model="MgForm.Credit" type="number" placeholder="请输入课程学分"></el-input>
            </el-form-item>
            <el-form-item label="描述信息">
              <el-input v-model="MgForm.Content" type="textarea" placeholder="请输入描述信息"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="AddCourse()">提交</el-button>
            </el-form-item>
          </el-form>
        </el-dialog>

    <el-form :inline="true"  :model="msgfrom"   label-width="100px"  size="medium" class="manage-form" style="width:100%">
    <el-table :data="tableData" style="width:100%" >
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column type="index" width="50"></el-table-column>

            <el-table-column prop="Cno" label="课程号" width="10%" align="center"></el-table-column>

            <el-table-column prop="CourseName" label="课程名称" width="150%" align="center"></el-table-column>

            <!--<el-table-column prop="Courseteacher" label="任教老师" align="center"></el-table-column>-->
            <el-table-column prop="StartDate"  align="center" width="220%" label="开设时间"></el-table-column>

            <!--el-table-column prop="credit"  align="center"  label="课程学分"></el-table-column-->

            
            <el-table-column  prop="Students"  align="center" width="170%" label="导入学生信息">
              <template slot-scope="{row}">
                    <el-button type="success" size="small" round width="150%" @click="Openform(row)">导入学生信息</el-button>
                  </template>
            </el-table-column>
          

            <el-table-column  prop="Students"  align="center" width="170%" label="查看学生信息">
              <template slot-scope="{row}">
                    <el-button type="success" round size="small" @click="getStudentList(row)">查看学生信息</el-button>
                  </template>
            </el-table-column>

            <el-table-column  prop="Assignmentlist"  align="center" width="170%" label="作业列表">
              <template slot-scope="{row}">
                    <el-button type="primary" round size="small" @click="getAssignmentlist(row)">查看课程作业</el-button>
              </template>
            </el-table-column>

            <el-table-column prop="operate" align="center" label="操作" width="210%">
              <template slot-scope="{row}">             
                <el-row>                 
                    <el-button type="primary" size="small" @click="OpenformchangeC(row)" >修改</el-button>
                    <el-button type="danger" size="small" @click="DelCourse(row)">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
    </el-table>
    
    </el-form>
                <el-dialog title="修改课程信息" :visible.sync="changecourseVisible" width="50%">
                    <el-form ref="form" :model="MgForm" label-width="120px" class="form-container">
                      <el-form-item label="课程名称">
                        <el-input v-model="MgForm.courseName" placeholder="请输入课程名称"></el-input>
                      </el-form-item>
                      <el-form-item label="开始时间">
                        <el-date-picker v-model="MgForm.startDate" type="datetime" placeholder="选择日期时间"></el-date-picker>
                      </el-form-item>
                      <el-form-item label="课程学分">
                        <el-date-picker v-model="MgForm.credit" type="number" placeholder="课程学分"></el-date-picker>
                      </el-form-item>
                      <el-form-item label="描述信息">
                        <el-input v-model="MgForm.Content" type="textarea" placeholder="请输入描述信息"></el-input>
                      </el-form-item>
                      <el-form-item>
                        <el-button type="primary" @click="ChangeCourse()">提交</el-button>
                        <el-button type="primary" @click="changecourseVisible = false">取消</el-button>
                      </el-form-item>
                    </el-form>
                </el-dialog>

                <el-dialog title="导入课程学生" :visible.sync="importVisible" width="50%">
                  <el-form ref="importorm" :model="CourseMsg" :rules="importFormRules" label-width="130px">
                    <el-form-item label="课程名" prop="Course" >
                      <el-input v-model="CourseMsg.Coursename" ></el-input>
                    </el-form-item>
                    <el-form-item label="课程号" prop="Cno" >
                      <el-input v-model="CourseMsg.Cno"></el-input>
                    </el-form-item>
                    <el-form-item label="上传文件:" prop="file">
                      <el-upload
                        class="import-demo"
                        ref="import"
                        :before-upload="beforeUpload"
                        :http-request="httpRequest"                          
                        :limit="1">
                        <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
                        <div slot="tip" class="el-upload__tip">只能上传execl文件，且不超过8M</div>
                      </el-upload> 
                    </el-form-item> 
                      <el-form-item>
                        <el-button type="primary" @click="submitImportForm">开始导入</el-button>
                        <el-button type="info" @click="importVisible = false">关闭窗口</el-button>
                      </el-form-item>                   
                    </el-form>                    
                </el-dialog>

                <el-row>
                  
                </el-row>
                
  </div>
</template>

<script>
import axios from 'axios';
import globalurl from '../views/GlobalMe.vue'

export default {
  name:'TeacherPage',
  created() {
      this.firstselectAll();
      this.selectAll();
      
    },
  data() {
    return {
      importVisible:false,
      addcourse:false,
      changecourseVisible:false,
      addAssignmentVisible:false,
      changeAssignmentVisible:false,
      FormData:new FormData(),
      Tno:'',
      temp:{
        Tno:'',
      },
      fileList: [], // 上传的文件列表
      CourseMsg:{
        Cno:'',
        Coursename:'',
      },
      msgfrom: {
        Cno:'',
        CourseName: '',
        //Courseteacher: "",
        Assignmentlist:[],
        StartDate: '',
        Students:[]
        //可能这里需要修改  导入的是execl表
      },
      //FormData:n,
      tableData: [
      {
        Cno:'',
        CourseName: '',
        //Courseteacher: "",
        Assignmentlist:[],
        StartDate: '',
        Students:[]
      },
      {
        Cno:'',
        CourseName: '',
        //Courseteacher: "",
        Assignmentlist:[],
        StartDate: '',
        Students:[]
      }
      ],
      MgForm: {
        CourseName: '',
        StartDate: '',
        Credit:'',
        Content: '',
        Tno:'',
      }
    }
  },
  methods: {
    LoginOut(){
      window.sessionStorage.clear()
      this.$router.push('/login')
    },

    firstselectAll() {      //根据教师工号查询所有课程
      const User = JSON.parse(localStorage.getItem('user'))
      console.log(" 哈哈哈  ",User)
      const username = localStorage.getItem('username')
      this.Tno = username
      this.temp.Tno = username
      this.msgfrom.Courseteacher = username
      this.tableData = User
      this.msgfrom = User
      console.log(" 哈哈哈  ",this.tableData)
    },
    selectAll(){
      console.log("Tno",this.Tno)
      axios({
        method:'post',
        url:globalurl.BASEURL+'TeacherCourseQuery',
        data:this.temp.Tno
      }).then(res=> {
         this.tableData = res.data.data;
         //this.totalCount = res.data.totalCount;
         })
    },
    Openform(row){
      this.CourseMsg.Cno = row.Cno
      this.CourseMsg.Coursename=row.Coursename
      this.importVisible=true
    },
    OpenformchangeC(row){
      this.MgForm.courseName=row.courseName
      this.MgForm.startDate = row.startDate
      /*this.MgForm.credit = row.credit
      this.MgForm.description = row.description*/
      this.changecourseVisible = true;
    },
    httpRequest(option) {
        this.fileList.push(option)
      },
    beforeUpload(file) {
  //const isxls=file.type === 'application/vnd.ms-excel'
    //const isxls = file.type === 'application/pdf'
    let extension = file.name.substring(file.name.lastIndexOf('.')+1);
    
    const isLt10M = file.size / 1024 / 1024 < 10
    
    if (extension !== 'xls' && extension !=='xlsx') {
      this.$message.error('上传文件只能是 execl 格式!')
    }
    if (!isLt10M) {
      this.$message.error('上传问价大小不能超过 10MB!')
    }
    return extension && isLt10M
    },
    submitImportForm() {           //导入学生信息
    console.log("sac杀杀杀")
        const params = new FormData()
        this.fileList.forEach((x) => {
          params.append('file', x.file)
        });
        params.append('Cno',this.CourseMsg.Cno)
		//这里根据自己封装的axios来进行调用后端接口
        axios({
          method:'post',
          //url:'http://localhost:9999/upload',
          url:globalurl.BASEURL+'ImportStudentInfo',
          //更新地址
          data:params,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            this.submitVisible=false;
            alert("上传成功")
            //this.tableData.Status = true;
            this.importVisible=false;
          }else{
            alert("上传失败")
            this.submitVisible=false;
          }
        })          
      },    
      //路径没写
    getStudentList(row){      //查看该课程的所有学生
      console.log("水水水水",row)
      axios({
        method:'post',
        url:globalurl.BASEURL+'CourseStudentQuery',
        data:row.Cno
      }).then(res=>{
        if (res.data.succession === 'true'){
          console.log()
          console.log("sss水水",res.data.data)
          localStorage.setItem('studentlist',JSON.stringify(res.data.data))
          localStorage.setItem('Cno',row.Cno)
          localStorage.setItem('Coursename',row.Coursename)
          window.location.href = "/students";
          //window.location.href = "/teacher";
          }else if (res.data.data === null){
             alert("该课程还未导入学生，请先导入学生信息")
          }
      })
    },
    //查看课程作业
    getAssignmentlist(row){ 

      console.log("宋湛杰") 
      //window.location.href="/assignment"
      axios({
        method:'post',
        url:globalurl.BASEURL+'CourseQuery',
        //还需修改
        data:row.Cno   //根据课程号找课程作业
      }).then(res => res.data).then(res => {
        console.log(res)
        if(res.succession ==='true'){
          localStorage.setItem('assignment',JSON.stringify(res.data))
          localStorage.setItem('Cno',row.Cno)
          localStorage.setItem('Coursename',row.Coursename)
          console.log("r哈哈哈哈",JSON.stringify(res.data))
          window.location.href="/assignment"
        }else{
          alert("获取失败")
        }
      }) 
    },
    AddCourse(){
      
      this.MgForm.Tno=this.Tno
      console.log("mgform=",this.MgForm)
      axios({
        method:'post',
        url:globalurl.BASEURL+'CourseAdd',
        data:this.MgForm,
        Headers:{"Content-Type":"multipart/FormData" },
      }).then(res=>res.data).then(res=>{
        if (res.succession === 'true'){
          this.selectAll();
          this.addcourse = false;
                alert("添加成功")
            }else{
                alert("添加失败")
            }
      })
    },
    ChangeCourse(){
      axios({
        method:'post',
        url:'',
        //url:globalurl.BASEURL+'do_login',
        data:this.MgForm
      }).then(res=>{
        if (res.data.data === 'true'){
          this.selectAll();
          alert("修改成功")
          }else {
             alert("修改失败")
          }
      })
      //暂时为实现
    },
    addAssignment(){
      this.addAssignmentVisible=true;

    },
    submitForm() {
      // 在这里提交表单数据，可以通过 API 发送请求给后端
      console.log(this.form)
    },
    DelCourse(row){
      const Data = {
        Tno:this.Tno,
        Cno:row.Cno,
      }
      console.log(Data)
      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
      }).then(() => {
        axios({
                method:"post",
                url:globalurl.BASEURL+'CourseDelete',
                //url:"http://localhost:8080/TeacherDelete",
                //data:row.Cno,
                data:Data,
            }).then(res=>{
              console.log("af大",res.data)
                if(res.data.succession==='true'){                        
                    this.selectAll();
                    
                    this.$message({
                      message: '已成功删除',
                        type: 'success'
                    });
                    this.selectAll();
                  }
                })
            }).catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消删除'
                });
            });
    },    
  }
}
</script>


<style scoped>
.container {
  max-width: 100%;
  
  margin: 0 auto;
  padding: 295px;
  background-image: url('../assets/images/background1.jpg');
  background-size: cover;
  color: #fff;
  font-family: "Microsoft Yahei";
}

.title {
  font-size: 32px;
  color: #fff;
  margin-bottom: 20px;
  text-shadow: 0 2px 2px rgba(0, 0, 0, 0.3);
}

.form-container {
  border: 1px solid rgba(255, 255, 255, 0.3);
  padding: 100px;
  border-radius: 100px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  background-color: rgba(255, 255, 255, 0.2);
}

.submit-button {
  width: 200px;
  background-color: #fff;
  color: #333;
  border: none;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.3);
}

.submit-button:hover {
  background-color: #f0f0f0;
  color: #333;
}
</style>