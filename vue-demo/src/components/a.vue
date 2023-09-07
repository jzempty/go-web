<template>

  
    <el-form 
    :inline="true" 
    :model="homeworkform" 
    label-width="100px" 
    ref="form"
    margin="0 auto"
    class="homework-form">
      <div id="homeWork"><h1>{{this.Course}}的作业</h1></div>
  
      
      
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
  
  
            <el-table-column prop="Sname" label="学生姓名" align="center"></el-table-column>
    
            <el-table-column prop="Status" align="center" label="提交状态">
              <div>
                {{Status ? '已提交' : '未提交'}}
              </div>
            </el-table-column>
  
            <el-table-column prop="Scope" label="作业得分" align="center"></el-table-column>
  
            <el-table-column prop="credit" label="评分" align="center"> 
              <template slot="label">操作</template>
              <el-row>
                <el-button type="primary" plain @click="scopeVisible=true">作业评分</el-button>
              </el-row>
  
            <el-dialog title="作业评分" :visible.sync="scopeVisible" width="50%">
              <el-form ref="form" :model="homeworkform" label-width="120px" class="form-container">
                  <el-form-item label="作业分数">
                    <el-input v-model="homeworkform.credit" placeholder="请输入作业分数"></el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="setScope">提交</el-button>
                    <el-button type="info" @click="scopeVisible=false">取消</el-button>
                  </el-form-item>
  
              </el-form>
              </el-dialog>
  
              
            </el-table-column>
            
              
  
            <el-table-column prop="report" align="center" label="操作">
              
                    <template slot-scope="scope">
                      <el-button
                        type="primary"
                        @click="downloadReport(scope.row)">
                        下载报告
                      </el-button>
                    </template>
            </el-table-column>
        </el-table>
        
     </el-form>
  
    
  </template>
  
  <script>
  
  
  
  export default {
    name:'homeWork',
    /*created() {
        const User = JSON.parse(localStorage.getItem('user'))
        console.log(" 哈哈哈  ",User)
        this.tableData = User
        this.user = User
        console.log(" 哈哈哈  ",this.tableData)
      },*/
    data() {
      return {
        dialogVisible:false,
        scopeVisible:false,
        //rate:'',
        Rate:{
          rate:'',
        },
        texts:['不及格','60-70','70-80','80-90','>90'],
        //Credit:0,
        homeworkform:{
            Sname:'',   //学生名
            credit:'',     //学号
            Scope:'',
            Status:'',  //作业提交情况
            reports:[{}],
        },
        FormData:new FormData(),
        tableData:[{
            Sname:'szj',
            credit:'',      
            Scope:'a',      
            Status:true,
            reports: [
              {
              name: '张三',
              reportUrl: 'http://www.example.com/report1.pdf'
              }]
            },
              {
              Sname:'szj',
              credit:'',     
              Scope:'a',       
              Status:true,
              reports: [{
              name: '张三',
              reportUrl: 'http://www.example.com/report1.pdf'
            }],
        }],
      };
    }, 
    methods: {  
      showDialog(rowData) {
        // 打开弹窗
        this.dialogVisible = true;
        // 传递参数
        console.log(rowData);
      },
      setScope(){
        this.scopeVisible = true;
      },
      rateChange(value) {
      this.row.Credit = value; // 将评分数值赋值给变量rate
      console.log(this.row.Credit)
  
      
    },        
  }
  };
  </script>
  
  