<template>
    <el-form 
    :inline="true" 
    :model="assignment" 
    label-width="100px"
    ref="form"
    size="medium"
    margin="0 auto"
    class="assign-form">
      <div id="AssignM"><h1>{{this.tableData.Sname}}作业</h1></div>
        <el-row>
          <el-button type="primary" plain @click="addAssignmentVisible=true">添加作业</el-button>      
        </el-row>
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
  
  
            <el-table-column prop="Title" label="作业名" align="center"></el-table-column>
  
            <el-table-column prop="Deadline" label="截止时间" align="center"></el-table-column>
  
            <el-table-column prop="Status" align="center" label="查看提交情况">
              <template slot-scope="{row}">
              <el-button type="success" @click="checksubmit(row)">查看提交情况</el-button>
              </template>
            </el-table-column>
  
            <el-table-column align="center" label="操作"> 
              <template slot-scope="{row}">               
              <el-button size="small" type="danger" @click="DelAssign(row)">删除作业</el-button>       
              <el-button size="small" type="primary" @click="Openform(row)">修改作业</el-button> 
              </template>           
            </el-table-column>
        </el-table>
        
        <el-dialog title="添加作业" :visible.sync="addAssignmentVisible" width="50%">
          <el-form ref="form" :model="assignment" label-width="120px" class="form-container">

            <el-form-item label="作业名称">
              <el-input v-model="assignment.Title" placeholder="请输入作业名称"></el-input>
            </el-form-item>

            <el-form-item label="截止时间">
              <el-date-picker v-model="assignment.Deadline" type="datetime" placeholder="选择日期时间"></el-date-picker>
            </el-form-item>

            <el-form-item label="描述信息">
              <el-input v-model="assignment.Context" type="textarea" placeholder="请输入描述信息"></el-input>
            </el-form-item>

          <el-row>
            <el-form-item>
              <el-button type="primary" @click="addAssignmentinCourse()" >提交</el-button>
              <el-button type="primary" @click="addAssignmentVisible = false" >取消</el-button>
            </el-form-item>
          </el-row>
            
          </el-form>
        </el-dialog>
        <el-dialog title="修改作业" :visible.sync="changeAssignmentVisible" width="50%">
          <el-form ref="form" :model="assignment" label-width="120px" class="form-container">

            <el-form-item label="作业名称">
              <el-input v-model="assignment.Title" placeholder="请输入作业名称"></el-input>
            </el-form-item>

            <el-form-item label="截止时间">
              <el-date-picker v-model="assignment.Deadline" type="datetime" placeholder="选择日期时间"></el-date-picker>
            </el-form-item>

            <el-form-item label="描述信息">
              <el-input v-model="assignment.Context" type="textarea" placeholder="请输入描述信息"></el-input>
            </el-form-item>

          <el-row>
            <el-form-item>
              <el-button type="primary" @click="changeAssign()" >提交</el-button>
              <el-button type="primary" @click="changeAssignmentVisible = false" >取消</el-button>
            </el-form-item>
          </el-row>
            
          </el-form>
        </el-dialog>

     </el-form>
  
    
  </template>
  
  <script>
  import axios from 'axios';
  import globalurl from '../views/GlobalMe.vue'

  export default {
    name:'AssigM',
    created() {
      const cno = localStorage.getItem('Cno')

      this.Cno = cno
      console.log("sa",this.Cno)
       this.selectAll();
      },
    data() {
      return {
        addAssignmentVisible:false,
        changeAssignmentVisible:false,
        Cname:'',
        Cno:'',
        assignment:{
            Title:'',
            Deadline:'',
            //所有的与时间有关的都要修改   +入时间函数
            Status:'',
            Context:'',
            Cno:'',
        },
        Assignment:{
            Title:'',
            Deadline:'',
            //所有的与时间有关的都要修改   +入时间函数
            Status:'',
            Context:'',
            Cno:'',
        },
        FormData:new FormData(),
        tableData:[{
            Title:'实验一',
            Deadline:'2023-06-01',
            Status:'',
            Cno:'',
        }],
      };
    }, 
    methods: {  
      selectAll(){
        
        /*const Assign = JSON.parse(localStorage.getItem('assignment'))
        this.Cname = localStorage.getItem('Coursename')
        this.Cno=localStorage.getItem('Cno')
        this.tableData = Assign */       
        axios({
          method:'post',
          url:globalurl.BASEURL+'CourseQuery',
          data:this.Cno,
        }).then(res=> {
              this.tableData = res.data.data;
              //this.totalCount = res.data.totalCount;
        })
      },
      Openform(row){
        this.assignment.Title = row.Title
        this.assignment.Deadline = row.Deadline
        this.assignment.Context = row.Context
        this.assignment.Cno = this.Cno
        this.changeAssignmentVisible=true
      },
    checksubmit(row){
      const params = new FormData()
      params.append('Cno',this.Cno)
      params.append('Title',row.Title)
      console.log("啦啦啦",params),
      axios({
        method:'post',
        url:globalurl.BASEURL+'AssignmentConditionQuery',
        data:params,
      }).then(res=>{
        if (res.data.succession === 'true'){
          //this.$message
          alert("即将抵达")
          localStorage.setItem('studentassignment',JSON.stringify(res.data.data))
          localStorage.setItem('Cno',this.Cno)
          localStorage.setItem('Title',row.Title)
          console.log("还好",res.data.data)
          window.location.href = '/studentassign';
        }else{
          alert("跳转失败")
        }
      })

        
    },
    addAssignmentinCourse(){
      const params = new FormData();
        params.append("Cno",this.Cno)
        params.append("Title",this.assignment.Title)
        params.append("Deadline",this.assignment.Deadline)
        console.log("时间",)
        params.append("Context",this.assignment.Context)
        this.assignment.Cno = this.Cno
        axios({
          method:'post',
          url:globalurl.BASEURL+'AssignmentAdd',
          data:this.assignment,
          //Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
            if (res.succession === 'true'){
              this.selectAll();
              this.addAssignmentVisible = false,
                alert("添加成功")
                //还有为实现的刷新表格
            }else{
                alert("添加失败")
            }
        })
    },
    DelAssign(row){
        /**this.FormData.append(Cname)
        this.FormData.append(Aname) */
        this.Assignment.Cno = this.Cno
        this.Assignment.Title = row.Title
        axios({
            method:'post',
            url:globalurl.BASEURL+'AssignmentDelete',
            data:this.Assignment
        }).then(res=>res.data).then(res=>{
            if (res.succession === 'true'){
              this.selectAll();
                alert("删除成功")
            }else{
                alert("删除失败")
            }
        })
    },
    changeAssign(){              //修改细节   
        axios({
            method:'post',
            url:globalurl.BASEURL+'AssignmentUpdate',
            data:this.assignment
        }).then(res=>res.data).then(res=>{
            if (res.succession === 'true'){
              this.selectAll();
                alert("修改成功")
            }else{
                alert("修改失败")
            }
        })
    },    
  }
}
  </script>
 <script setup>

</script> 
  