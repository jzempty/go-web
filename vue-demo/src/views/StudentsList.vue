<!--某一门课程的学生信息  -->

<template>
    <div id="CourseStudents" class="container">
      <div><h1>{{this.Coursename}}学生信息表</h1></div> 

      <el-row>
        <el-button type="primary" plain @click="addstudentVisible = true">添加学生</el-button> 
      </el-row>
      <br> 
        <el-dialog title="添加课程学生" :visible.sync="addstudentVisible" width="30%">
          <el-form ref="form" :model="studentsmsg" label-width="30%" class="form-container">

            <el-form-item label="学生姓名">   
              <el-input v-model="studentsmsg.Sname" placeholder="请输入学生姓名"></el-input>
            </el-form-item>

            <el-form-item label="学生学号" width="10%">
              <el-input v-model="studentsmsg.Sno" width="15%" placeholder="请输入学生学号"></el-input>
            </el-form-item>
            <el-form-item label="学生密码" width="10px">
              <el-input v-model="studentsmsg.Password" width="5%" placeholder="请输入学生密码"></el-input>
            </el-form-item>

            <el-form-item label="班级" width="10px">
              <el-input v-model="studentsmsg.ClassName" width="5%" placeholder="请输入班级"></el-input>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" @click="addstudent()">提交</el-button>
              <el-button type="primary" @click="addstudentVisible=false">取消</el-button>
            </el-form-item>
          </el-form>
        </el-dialog>
        

        <!--搜索表单-->
        <span><el-form :inline="true" :model="searchstudent" class="demo-form-inline">
            <el-form-item label="姓名">
                <el-input v-model="searchstudent.Sname" placeholder="学生姓名"></el-input>
            </el-form-item>           
            <el-form-item label="班级">
                <el-radio-group v-model="searchstudent.ClassName">
                    <el-radio label="信计201"></el-radio>
                    <el-radio label="信计202"></el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="Search">查询</el-button>
            </el-form-item>
        </el-form></span>

        <el-row>
          <el-button type="primary" plain @click="selectAll()">查询所有学生</el-button> 
        </el-row>

        

      <el-form 
      :inline="true" 
      :model="studentsMsg" 
      label-width="100px" 
      margin="0 auto"
      class="students-form">      
           <el-table :data="tableData" style="width: 100%"  stripe>
              <el-table-column type="selection" width="55"> </el-table-column>
              <el-table-column type="index" width="50"></el-table-column>
       
              <el-table-column prop="Sname" label="学生姓名" align="center"></el-table-column>
      
              <el-table-column prop="Sno" label="学生学号" align="center"></el-table-column>
  
              <el-table-column prop="ClassName" align="center" label="所在班级"></el-table-column>
              <el-table-column align="center" label="操作">
            <template slot-scope="{row}">          
              <!--<el-button  type="primary"  @click="Openform(row)">修改学生信息</el-button>-->
              <el-button  type="danger" class="el-icon-delete" @click="Delstudent(row)">删除学生</el-button>          
            </template>
              </el-table-column>
          </el-table>        
       </el-form>
    
       <el-dialog title="修改学生信息" :visible.sync="changestudentVisible" width="50%">
            <el-form ref="form" :model="studentsMsg" label-width="120px" class="form-container">
                <el-form-item label="学生姓名">
                    <el-input v-model="studentsMsg.Sname" placeholder="请输入学生姓名"></el-input>
                </el-form-item>

                <el-form-item label="学生学号">
                    <el-input v-model="studentsMsg.Sno" placeholder="请输入学生学号"></el-input>
                </el-form-item>

                <el-form-item label="所属班级">
                    <el-input v-model="studentsMsg.Classname" placeholder="请输入学生所属班级"></el-input>
                </el-form-item>
            <el-row>
              <el-form-item>                  
                <el-button type="primary" @click="submitchangeMsg()">提交</el-button>
                <el-button type="primary" @click="changestudentVisible= false">取消</el-button>
              </el-form-item>
            </el-row>
            </el-form>
        </el-dialog>
        
  
      </div> 
    </template>
    
    <script>
 // import axios from 'axios';

import axios from 'axios';
import globalurl from '../views/GlobalMe.vue'
  
    export default {
      name:'CourseStudents',
      created() {
          this.firstselectAll();
          this.selectAll();
        },
      data() {
        return {
          dialogVisible:false,
          addstudentVisible:false,
          changestudentVisible:false,
          Coursename:'',
          Cno:'',
          FormData:new FormData(),
          searchstudent:{
            Sname:'',
            ClassName:'',
          },
          studentsMsg:{
              Sname:'',             //学生名
              Sno:'',               //学号
              ClassName:'',          //班级   
              Cno:'',    
              Password:'',

          },  
          studentsmsg:{             //用于新增学生的结构体
              Sname:'',             //学生名 
              Sno:'',               //学号
              ClassName:'',          //班级   
              Cno:'',    
              Password:'',

          }, 
          delstudent:{
              Sname:'',             //学生名
              Sno:'',               //学号
              ClassName:'',          //班级   
              Cno:'',    
              Password:'',

          },
          students:[],        
          tableData:[
            {
            Sname:'szj',             //学生名
            Sno:'001',               //学号
            ClassName:'1'          //班级           
            },
            {
            Sname:'ljz',             //学生名
            Sno:'002',               //学号
            ClassName:'2'          //班级           
            }
            ], 
          rules: {
            Sname: [
              {required: true, message: "请输入姓名", trigger: "blur"},
              {min: 1, max: 20, message: '请勿超过20个字符！', trigger: 'blur'},
              {validator: (rule, value, callback) => {
              const reg = /^[\u4e00-\u9fa5]+$/  //正则表达式 只能输入英文、汉字与数字
              if (!reg.test(value)) {
                callback(new Error('请勿输入特殊字符'))
              } else {
                callback()
              }}}
            ],
            Sno: [
              {required: true, message: "请输入学生学号", trigger: "blur"},
              {min: 10, max: 10, message: '只能输入数字并且长度在10字符', trigger: 'blur'},
              {validator: (rule, value, callback) => {
              const reg = /^[0-9]+$/  //正则表达式 只能输入数字
              if (!reg.test(value)){
                callback(new Error('请勿输入特殊字符'))
              }else if(length(reg.test(value))!= 10){
                callback(new Error('输入学号不符合要求'))
              }else {
                callback()
              }}}
            ],
      },
        };
      }, 
      methods: { 
        firstselectAll(){
          const User = JSON.parse(localStorage.getItem('studentlist'))
          const Coursename = localStorage.getItem('Coursename')
          const Cno = localStorage.getItem('Cno')
          console.log(" 哈哈哈  ",User)
          console.log('sasa',Coursename)
          this.tableData = User
          this.studentsMsg = User  
          this.Coursename = Coursename 
          this.Cno = Cno
        },
        selectAll(){
          axios({
            method:'post',
            url:globalurl.BASEURL+'CourseStudentQuery',
            data:this.Cno
          }).then(res=> {
              this.tableData = res.data.data;
              //this.totalCount = res.data.totalCount;
                })
        },
        Openform(row){       //包装弹框表单的消息--和点击的那一行表格信息对应
            this.studentsMsg.Sname = row.Sname
            this.studentsMsg.Sno = row.Sname
            this.studentsMsg.ClassName = row.ClassName
            this.changestudentVisible=true
        }, 
        Search(){          //查询对应课程的所有学生
          this.students=[]
          for(let i=0;i<this.studentsMsg.length;i++){
            if(this.searchstudent.Sname.length>0){
              if (this.studentsMsg[i].ClassName === this.searchstudent.ClassName && this.studentsMsg[i].Sname==this.searchstudent.Sname){
                this.students.push(this.studentsMsg[i])
              }
              }else{
                if (this.studentsMsg[i].ClassName == this.searchstudent.ClassName){
                  this.students.push(this.studentsMsg[i])
               }
            }
            
          }
          this.tableData = this.students
        },
        addstudent(){       //添加课程新的学生   表单获取
          console.log("assa",this.studentsmsg)
          axios({
            method:'post',
            url:globalurl.BASEURL+'StudentAdd',
            data:this.studentsmsg,
          }).then(res=>{
              if (res.data.succession ==='true'){
                localStorage.setItem('studentassignment',res.data.data)
                this.selectAll;
                alert("添加成功")
              }else{
                alert("添加失败")
              }
            })
        },
        submitchangeMsg(){
            axios({
                method:'post',
                url:globalurl.BASEURL+'CourseStudentQuery',
                data:this.studentsMsg,
            }).then(res=>{
                if (res.data.succession === 'true'){
                    alert("aa")
                }else{
                    alert("bb")
                }
            })
        },
        Delstudent(row){
            this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
              console.log("dada")
              this.delstudent.Cno = this.Cno
              this.delstudent.Sno = row.Sno
              console.log("dada",this.studentsMsg)
                axios({
                    method:"post",
                    url:globalurl.BASEURL+'DeleteSchedule',   //删除某门课程里面的学生
                    data:this.delstudent,
                }).then(res=>{
                    if(res.data.succession==='true'){
                        
                       this.selectAll();
                        //提示已删除
                        this.$message({
                            message: '已成功删除',
                            type: 'success'
                        });
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
    };
  </script>
    
    