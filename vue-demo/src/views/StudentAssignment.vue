<!--某一门作业提交情况  -->

<template>
  
  <div id="homeWork" class="container">
    <el-header><h1>作业的提交情况</h1></el-header>
    

    <el-main>
     <!--搜索表单-->
     <el-form :inline="true" :model="searchStudentAssign" class="demo-form-inline">
            <el-form-item label="姓名">
                <el-input v-model="searchStudentAssign.Sname" placeholder="学生姓名"></el-input>
            </el-form-item>   
                    
            <el-form-item label="提交状态">
                <el-radio-group v-model="searchStudentAssign.Status">
                    <el-radio label="已提交"></el-radio>
                    <el-radio label="未提交"></el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" class="el-icon-view" @click="Search()">查询</el-button>
            </el-form-item>
        </el-form>

    <el-form 
    :inline="true" 
    :model="homeworkform" 
    label-width="100px" 
    margin="0 auto"
    class="homework-form">      
         <el-table :data="tableData" style="width: 100%"  stripe>
            <el-table-column type="selection" width="55"> </el-table-column>
            <el-table-column type="index" width="50"></el-table-column>
  
  
            <el-table-column prop="Sname" label="学生姓名" align="center"></el-table-column>
    
            <el-table-column prop="Sno" label="学生学号" align="center"></el-table-column>

            <el-table-column prop="Status" align="center" label="提交状态">
            </el-table-column>

            <el-table-column prop="Score" label="作业得分" align="center"></el-table-column>

            <el-table-column prop="Credit" label="评分" align="center"> 
              <template slot-scope="{row}">
                <el-row>        
                <el-button type="primary" plain @click="Openform(row)">作业评分</el-button>
                </el-row>
              </template>             
            </el-table-column>

            <el-table-column width="150px" align="center" label="操作">            
                    <template slot-scope="{row}">
                      <a href='downloadurl' download='文件名'>下载报告</a>
                      <el-button size="small" type="text" class="el-icon-download" @click="downloadReport(row)">下载报告</el-button>
                      <el-button size="small" type="danger" class="el-icon-delete-solid" @click="DelHandle(scope.row.id)">删除</el-button>
                    </template>
            </el-table-column>
        </el-table>        
     </el-form>
  
    </el-main>
      <el-dialog title="作业评分" :visible.sync="scoreVisible" width="50%">
          <el-form ref="form" :model="Rate" label-width="120px" class="form-container">
            <el-form-item label="作业分数">
              <el-input v-model="Rate.credit" placeholder="请输入作业分数"></el-input>
                </el-form-item>
          <el-row>
            <el-form-item>                  
              <el-button type="primary" @click="setScore()">提交</el-button>
              <el-button type="primary" @click="scoreVisible= false">取消</el-button>
            </el-form-item>
          </el-row>
          </el-form>
      </el-dialog>

    </div> 
  
  </template>
  
  <script>
import axios from 'axios';
import globalurl from '../views/GlobalMe.vue'

  export default {
    name:'homeWork',
    created() {
        this.firstselectAll();
        this.selectAll();
        
      },
    data() {
      return {
        dialogVisible:false,
        scoreVisible:false,
        Params:{
          Cno:'',
          title:'',
        },
        h:{},
        AID:'',
        //rate:'',
        Rate:{
          credit:'',
        },
        texts:['不及格','60-70','70-80','80-90','>90'],
        //Credit:0,
        searchStudentAssign:{
          Sname:'',
          Sno:'',
          Status:'',
        },
        homeworkform:{
            Sname:'',   //学生名
            Sno:'',
            Credit:'',     //学号
            Score:'',
            Status:'',  //作业提交情况
            reports:{},
            Aid:'',
        },
        Status:'',
        FormData:new FormData(),
        Form:[],
        tableData:[{
            Sname:'szj',
            Sno:'',
            Credit:'',      
            Score:'a',      
            Status:true,
            reports:{},
            Aid:'',
            },
              {
              Sname:'szj',
              Sno:'',
              Credit:'',     
              Score:'a',       
              Status:'',
              reports:{},
              Aid:'',
            }
          ],

      };
    }, 
    methods: { 
      downloadReport(row){         //下载报告
        /*axios({
          method:'get',
          url:
          headers: {
            'Content-Type': 'application/json; charset=utf-8'
          },
        responseType: 'blob',
        }).then(res=>{
          console.log(res.data)
          //var content = res.headers['content-disposition'];
          //var name = content && content.split(';')[1].split('filename=')[1];
          //var fileName = decodeURIComponent(name)
          //downloadFile(res.data,fileName)
          })*/

      let xhr = new XMLHttpRequest();
      xhr.open('GET', globalurl.BASEURL+'GetFile?Aid='+row.Aid+'&Sno='+row.Sno, true);
      xhr.responseType = 'blob';

      xhr.onload = () => {
        if (xhr.status === 200) {
          let blob = new Blob([xhr.response], { type: 'application/pdf' });
          let url = window.URL.createObjectURL(blob);

          let link = document.createElement('a');
          link.href = url;
          link.download = `${row.title}实验报告.pdf`;
          link.click();

          window.URL.revokeObjectURL(url);
        }
      };
      xhr.send();
    }, 
      firstselectAll(){
        const p = localStorage.getItem('params')
        this.Params = p
        for (let i = 0;i<this.homeworkform.length; i++) {
        if(this.tableData[i].Status===true){
        this.tableData[i].Status='已提交'
        }else{
        this.tableData[i].Status='未提交'
        }
        }
      },
      selectAll(){
      const params = new FormData();
      params.append('Cno',localStorage.getItem('Cno'))
      params.append('Title',localStorage.getItem('Title'))
      axios({
        method:'post',
        url:globalurl.BASEURL+'AssignmentConditionQuery',
        data:params
      }).then(res=> {
         this.tableData = res.data.data;
         this.AID = this.tableData[0].Aid
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
        this.homeworkform=res.data.data
         //this.totalCount = res.data.totalCount;
         })
    },
    Openform(row){
      const h ={
        Aid:this.AID,
        Sno:row.Sno,
        Score:'',
      }
      this.h = h
      if (row.Status ==='未提交'){
        alert("学生未提交作业，请先提醒学生提交作业后在评分")
      }else{
        this.scoreVisible=true;
      }
      
    },
      showDialog(rowData) {
        // 打开弹窗
        this.dialogVisible = true;
        // 传递参数
        console.log(rowData);
      },
      setScore(){
        console.log(this.Rate.credit)
        this.h.Score = this.Rate.credit
        console.log("宋湛杰",this.h)
        axios({
          method:'post',
          url:globalurl.BASEURL+'AssignmentCorrect',
          data:this.h
        }).then(res=>res.data).then(res=>{
          if (res.succession === "true"){
            this.selectAll();
            alert("成功评分")
          }else{
            alert("评分失败")
          }
        })
      },
      Search(){             //查询某一次作业学生的提交情况  查询指定的信息
        this.Form=[]
      for(let i=0;i<this.homeworkform.length;i++){
      if(this.searchStudentAssign.Sname.length>0){
        if (this.homeworkform[i].Status == this.searchStudentAssign.Status && this.homeworkform[i].Sname==this.searchStudentAssign.Sname){
        this.Form.push(this.homeworkform[i])
        }
      }else{
        if (this.homeworkform[i].Status == this.searchStudentAssign.Status){
          this.Form.push(this.homeworkform[i])
        }
      }
      
     }
     this.tableData = this.Form
      },
      rateChange(value) {
      this.row.Credit = value; // 将评分数值赋值给变量rate
      console.log(this.row.Credit)

      
    }, 
    
    // 删除
    DelHandle(row) {                //删除某个学生的作业
        this.$confirm('是否确认删除该附件', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then( ()=>{
          const params = new this.FormData()
          params.append()
          axios({
            method:"post",
            //url:"http://localhost:8080/TeacherDelete",
            data:row,
          }).then(res=>{
                if(res.data=="success"){
                //提示已删除
                  this.$message({
                    message: '已成功删除',
                    type: 'success'
                  });
                }
              })
        })
      },
    }      
  };
</script>
  
  