<template>
    <div>
           
        <el-container style="height: 100%; border: 1px solid #eee">
            <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
                <el-menu :default-openeds="['1', '3']">
                <el-submenu index="1">
                    <template slot="title"><i class="el-icon-s-custom">教师</i></template>
                    <el-menu-item-group>                       
                    <el-menu-item index="1-1">
                        <el-button type="primary" size="small"  @click="selectAllTeacher()">查看教师信息</el-button>
                    </el-menu-item>  
                        <!--<el-menu-item index="1-2">
                            <el-button type="primary" size="small" @click="teacherVisible=true">导入新教师</el-button>
                        </el-menu-item> -->                     
                    </el-menu-item-group>                       
                </el-submenu>
                <!--<el-submenu index="2">
                    <template slot="title"><i class="el-icon-menu"></i>班级</template>
                        <el-menu-item-group>
                        <template slot="title">添加班级</template>
                        <el-menu-item index="2-1">
                            <el-button type="success" size="small" @click="studentVisible=true">添加新班级</el-button>
                        </el-menu-item>
                        </el-menu-item-group>
                    <el-menu-item-group>
                    <template slot="title">信计</template>
                    <el-menu-item index="2-1">
                        <template slot-scope="{row}">
                            <el-button type="primary" size="small" @click="selectAllClassStudents(row)">信计201</el-button>
                        </template>
                    </el-menu-item>
                    <el-menu-item index="2-2">
                        <template slot-scope="{row}">
                            <el-button type="primary" size="small" @click="selectAllClassStudents(row)">信计201</el-button>
                        </template>
                    </el-menu-item>
                    </el-menu-item-group>                        
                </el-submenu>-->
                <el-submenu index="2">
                    <template slot="title"><i class="el-icon-menu">班级</i></template>
                    <el-menu-item-group>                       
                    <el-menu-item index="2-1">
                        <el-button type="primary" size="small"  @click="selectAllClass()">查看班级信息</el-button>
                    </el-menu-item>  
                        <!--<el-menu-item index="1-2">
                            <el-button type="primary" size="small" @click="teacherVisible=true">导入新教师</el-button>
                        </el-menu-item> -->                     
                    </el-menu-item-group>                       
                </el-submenu>
                <el-submenu index="3">
                    <template slot="title"><i class="el-icon-s-custom">学生</i></template>
                    <el-menu-item-group>                       
                    <el-menu-item index="3-1">
                        <el-button type="primary" size="small"  @click="selectAllStudent()">查看学生信息</el-button>
                    </el-menu-item>                       
                    </el-menu-item-group>                       
                </el-submenu>
                <el-submenu index="4">
                    <template slot="title"><i class="el-icon-s-grid"></i>课程</template>
                    <el-menu-item-group>
                    <el-menu-item index="4-1">
                        <el-button type="primary" size="small" @click="selectAllCourses()">查看课程信息</el-button>
                    </el-menu-item>
                    </el-menu-item-group>
                </el-submenu>
                </el-menu>
            </el-aside>
                
                <el-container>
                    <el-header style="text-align: center; font-size: 12px ">
                    <el-dropdown>
                        <i class="el-icon-s-opportunity" style="margin-right: 15px"></i>
                        <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item>查看</el-dropdown-item>
                        <el-dropdown-item>新增</el-dropdown-item>
                        <el-dropdown-item>删除</el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                    <span color="red">欢迎您:szj 管理员</span><br>
                    <span color="red">欢迎您:szj 管理员</span>
                    
                    </el-header>
                    
                    <el-main>
                    <!--<el-table :data="tableData">
                        <el-table-column prop="date" label="日期" width="140">
                        </el-table-column>
                        <el-table-column prop="name" label="姓名" width="120">
                        </el-table-column>
                        <el-table-column prop="address" label="地址">
                        </el-table-column>
                    </el-table>-->
 
                    <el-row>
                        <el-button type="primary" plain @click="addcourseVisible=true">新建课程</el-button> 
                    </el-row>
                    <el-row>
                        <br>
                    </el-row>
                    <el-row>
                        <el-button type="primary" plain @click="addclassVisible=true">添加班级</el-button>
                    </el-row>
                    <br>
                    <el-row>
                        <el-button type="primary" plain @click="importteacherVisible=true">导入新教师</el-button>                    
                    </el-row>

                    <el-dialog title="导入教师信息" :visible.sync="importteacherVisible" width="30%">
                            <!--<el-form 
                            :inline="true" 
                            :model="teachers" 
                            label-width="100px" 
                            margin="0 auto"
                            class="students-form">      
                                <el-table :data="Ttabledata" style="width: 100%"  stripe>
                                    <el-table-column type="selection" width="55"> </el-table-column>
                                    <el-table-column type="index" width="50"></el-table-column>
                            
                                    <el-table-column prop="Tno" label="教师工号" align="center"></el-table-column>

                                    <el-table-column prop="Tname" label="教师姓名" align="center"></el-table-column>                                    
                        
                                    <el-table-column prop="Tgender" align="center" label="所在班级"></el-table-column>

                                    <el-table-column prop="Tdept" label="学院" align="center"></el-table-column>

                                    <el-table-column align="center" label="操作">
                                    <template slot-scope="{row}">          
                                    <el-button  type="primary" size="small" @click="Openform(row)">修改学生信息</el-button>
                                    <el-button  type="danger" size="small" class="el-icon-delete" @click="Delstudent(row)">删除学生</el-button>          
                                    </template>
                                    </el-table-column>
                                </el-table>        
                            </el-form>-->
                        <el-row>
                            <el-button type="primary" @click="singleteacherVisible=true">导入单个</el-button>
                        </el-row>
                        <br>
                        <el-row>
                            <el-button type="primary" @click="fileteacherVisible=true">文件导入</el-button>
                        </el-row>
                        </el-dialog>

                        <el-dialog title="查看教师信息" :visible.sync="teacherVisible" width="80%">
                            <el-form 
                            :inline="true" 
                            :model="teachers" 
                            label-width="100px" 
                            margin="0 auto"
                            class="students-form">      
                                <el-table :data="Ttabledata" style="width: 100%"  stripe>
                                    <el-table-column type="selection" width="55"> </el-table-column>
                                    <el-table-column type="index" width="50"></el-table-column>
                            
                                    <el-table-column prop="Tno" label="教师工号" align="center"></el-table-column>

                                    <el-table-column prop="Tname" label="教师姓名" align="center"></el-table-column>                                    
                        
                                    <el-table-column prop="Tsex" align="center" label="所在班级"></el-table-column>

                                    <el-table-column prop="Tdept" label="学院" align="center"></el-table-column>

                                    <el-table-column align="center" label="操作">
                                    <template slot-scope="{row}">          
                                    <el-button  type="primary" size="small" icon="el-icon-edit" width="10px" @click="Openform(row)">修改教师信息</el-button>
                                    <el-button  type="danger" size="small" class="el-icon-delete" width="10px" @click="Delteacher(row)">删除教师  </el-button>          
                                    </template>
                                    </el-table-column>
                                </el-table>        
                            </el-form>
                        </el-dialog>

                        <el-dialog title="新增教师" :visible.sync="singleteacherVisible" width="30%">
                            <el-form ref="form" :model="teachers"  label-width="120px"  margin="0 auto" class="students-form">      
                                <el-form-item label="教师名称" width="30px">
                                    <el-input v-model="teachers.Tname" width="30" placeholder="请输入教师名字"></el-input>
                                </el-form-item>  
                                <el-form-item label="教师工号">
                                    <el-input v-model="teachers.Tno" placeholder="请输入教师工号"></el-input>
                                </el-form-item>
                                
                                <el-form-item label="学院">
                                    <el-input v-model="teachers.Tdept" placeholder="请输入教师所属学院"></el-input>
                                </el-form-item>   
                                <el-form-item label="性别" >
                                    <el-radio-group v-model="teachers.Tsex">
                                        <el-radio label="男"></el-radio>
                                        <el-radio label="女"></el-radio>
                                    </el-radio-group>
                                </el-form-item>  
                            <el-form-item>
                                <el-button type="primary" @click="AddTeacher()">提交</el-button>
                            </el-form-item>
                            </el-form>
                        </el-dialog>

                        <el-dialog title="导入教师" :visible.sync="fileteacherVisible" width="30%">
                        <el-form ref="importorm" :model="teachers" :rules="importFormRules" label-width="130px">
                            <el-form-item label="上传文件:" prop="file" align="center">
                            <el-upload
                                class="import-demo"
                                ref="import"
                                :before-upload="beforeUpload"
                                :http-request="httpRequest"                          
                                :limit="1">
                                <el-button slot="trigger" size="small" type="primary" align="center">选取文件</el-button>
                                <div slot="tip" class="el-upload__tip" align="center">只能上传execl文件，且不超过8M</div>
                            </el-upload> 
                            </el-form-item> 
                            <el-form-item>
                                <el-button type="primary" @click="submitImportForm">开始导入</el-button>
                                <el-button type="info" @click="fileteacherVisible = false">关闭窗口</el-button>
                            </el-form-item>                   
                            </el-form>                    
                        </el-dialog>

                        <el-dialog title="查看课程信息" :visible.sync="courseVisible" width="80%">
                            <el-form 
                            :inline="true" 
                            :model="courses" 
                            label-width="100px" 
                            margin="0 auto"
                            class="students-form">      
                                <el-table :data="Ctabledata" style="width: 100%"  stripe>
                                    <el-table-column type="selection" width="55"> </el-table-column>
                                    <el-table-column type="index" width="50"></el-table-column>
                            
                                    <el-table-column prop="Cno" label="课程号" align="center"></el-table-column>

                                    <el-table-column prop="CourseName" label="课程名称" align="center"></el-table-column>                                    
                        
                                    <el-table-column prop="Credit" align="center" label="课程学分"></el-table-column>

                                    <el-table-column prop="Tno" align="center" label="任教教师"></el-table-column>

                                    <el-table-column prop="StartDate" align="center" label="课程开设时间"></el-table-column>

                                    <el-table-column align="center" label="操作">
                                    <template slot-scope="{row}">          
                                    <el-button  type="primary" size="small" icon="el-icon-edit" @click="Openform(row)">修改课程信息</el-button>
                                    <el-button  type="danger" size="small" class="el-icon-delete" @click="Delcourse(row)">删除课程  </el-button>          
                                    </template>
                                    </el-table-column>
                                </el-table>        
                            </el-form>
                        </el-dialog>

                        <el-dialog title="新增课程" :visible.sync="addcourseVisible" width="30%">
                            <el-form ref="form" :model="courses"  label-width="120px"  margin="0 auto" class="students-form">      
                                <el-form-item label="课程名称" width="30px">
                                    <el-input v-model="courses.CourseName" width="30" placeholder="请输入课程名称"></el-input>
                                </el-form-item>  
                                <el-form-item label="课程学分">
                                    <el-input v-model="courses.Credit" placeholder="请输入课程学分"></el-input>
                                </el-form-item>  
                                <el-form-item label="课程教师工号">
                                    <el-input v-model="courses.Tno" placeholder="请输入课程教师工号"></el-input>
                                </el-form-item>
                                <el-form-item label="课程开始时间">
                                    <el-date-picker v-model="courses.StartDate" type="datetime" placeholder="选择日期时间"></el-date-picker>
                                </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="AddCourse()">提交</el-button>
                            </el-form-item>
                            </el-form>
                        </el-dialog>

                        <el-dialog title="新增班级" :visible.sync="addclassVisible" width="30%">
                            <el-form ref="form" :model="Class"  label-width="120px"  margin="0 auto" class="class-form">      
                                <el-form-item label="班级名称" width="30px">
                                    <el-input v-model="Class.ClassName" width="30" placeholder="请输入班级名称"></el-input>
                                </el-form-item>  
                                <el-form-item label="学院">
                                    <el-input v-model="Class.Cdept" placeholder="请输入班级所属学院"></el-input>
                                </el-form-item>
                                <el-form-item label="班主任">
                                    <el-input v-model="Class.PrimaryTeacher" placeholder="请输入班级班主任"></el-input>
                                </el-form-item>    
                            <el-form-item>
                                <el-button type="primary" @click="AddClass()">提交</el-button>
                            </el-form-item>
                            </el-form>
                        </el-dialog>

                        <el-dialog title="查看班级信息" :visible.sync="classVisible" width="80%">
                            <el-form 
                            :inline="true" 
                            :model="Class" 
                            label-width="100px" 
                            margin="0 auto"
                            class="students-form">      
                                <el-table :data="Classtabledata" style="width: 100%"  stripe>
                                    <el-table-column type="selection" width="55"> </el-table-column>
                                    <el-table-column type="index" width="50"></el-table-column>
                            
                                    <el-table-column prop="ClassId" label="班级号" align="center"></el-table-column>

                                    <el-table-column prop="ClassName" label="班级名称" align="center"></el-table-column>                                    
                        
                                    <el-table-column prop="Cdept" align="center" label="所属学院"></el-table-column>

                                    <el-table-column prop="PrimaryTeacher" align="center" label="班主任"></el-table-column>

                                    <el-table-column align="center" label="操作">
                                    <template slot-scope="{row}">          
                                    <el-button  type="primary" size="small" icon="el-icon-edit" @click="Openform(row)">修改班级信息</el-button>
                                    <el-button  type="danger" size="small" class="el-icon-delete" @click="Delclass(row)">删除班级  </el-button>          
                                    </template>
                                    </el-table-column>
                                </el-table>        
                            </el-form>
                        </el-dialog>

                        <el-dialog title="查看学生信息" :visible.sync="studentVisible" width="80%">
                            <el-form 
                            :inline="true" 
                            :model="students" 
                            label-width="100px" 
                            margin="0 auto"
                            class="students-form">      
                                <el-table :data="Stabledata" style="width: 100%"  stripe>
                                    <el-table-column type="selection" width="55"> </el-table-column>
                                    <el-table-column type="index" width="50"></el-table-column>
                            
                                    <el-table-column prop="Sno" label="学生学号" align="center"></el-table-column>

                                    <el-table-column prop="Sname" label="学生姓名" align="center"></el-table-column>                                    
                        
                                    <el-table-column prop="ClassName" align="center" label="所在班级"></el-table-column>

                                    <el-table-column align="center" label="操作">
                                    <template slot-scope="{row}">          
                                    <el-button  type="primary" size="small" icon="el-icon-edit" @click="Openfrom(row)">修改学生信息</el-button>
                                    <el-button  type="danger" size="small" class="el-icon-delete" @click="Delstudent(row)">删除学生  </el-button>          
                                    </template>
                                    </el-table-column>
                                </el-table>        
                            </el-form>
                        </el-dialog>

                        <el-dialog title="修改学生信息" :visible.sync="changestudentVisible" width="80%">
                            <el-form ref="importorm" :model="students" :rules="importFormRules" label-width="130px">
                                <el-form-item label="学生姓名" prop="Sname" >
                                <el-input v-model="students.Sname" ></el-input>
                                </el-form-item>
                                <!--<el-form-item label="学生学号" prop="Sno" >
                                <el-input v-model="students.Sno"></el-input>
                                <span>学生学号不能够修改 严重错误请删除学生重新添加</span>
                                </el-form-item>--> 
                                <el-form-item label="班级" prop="ClassName" >
                                <el-input v-model="students.ClassName"></el-input>
                                </el-form-item> 
                                <el-form-item>
                                    <el-button type="primary" @click="ChangeStudent()">提交</el-button>
                                    <el-button type="primary" @click="changestudentVisible = false">取消</el-button>
                                </el-form-item>                 
                            </el-form>  
                        </el-dialog>
                        

                    </el-main>
                    <el-row >
                            <el-button align="rigth" type="primary" class="el-icon-back" plain @click="LoginOut()">安全退出</el-button>          
                        </el-row>
                </el-container>
            </el-container>
</div>
</template>

<script>
import axios from 'axios';
import globalurl from '../views/GlobalMe.vue'
  export default {
    data() {
      return {
        importteacherVisible:false,
        fileteacherVisible:false,
        singleteacherVisible:false,
       teacherVisible:false,
       classVisible:false,
       studentVisible:false,
       changestudentVisible:false,
       addclassVisible:false,
       addcourseVisible:false,
       courseVisible:false,
       fileList: [], // 上传的文件列表
       teachers:{
        Tname:'',
        Tno:'',
        Tsex:'',
        Tdept:'',
        Firstpassword:''
       },
       Ttabledata:[{
        Tname:'lzw',
        Tno:'001',
        Tgender:'男',
        Tdept:'信计',
       }],
       students:{
        Sno:'',
        Sname:'',
        ClassName:'',
       },
       Stabledata:[
        {
            Sno:'11',
            Sname:'szj',
            ClassName:'信计202',
        },
       ],
       Class:{
        ClassName:'',
        ClassId:'',
        Cdept:'',
        PrimaryTeacher:''
       },
       Ctabledata:[{
        ClassName:'',
        ClassId:'',
        Cdept:'',
        PrimaryTeacher:''
       }],
       courses:{
        CourseName:'',
        Cno:'',
        StartDate:'',
        Credit:'',
        Tno:'',
       },
       Classtabledata:[
        {
            CourseName:'a',
            Cno:'1',
            Credit:'6',
            StartDate:'',
            Tno:'',
        }
       ],
       class:{     //添加班级
        Cdept:'',
        Classname:'',
        primary_teacher:''
       }
      }
    },
    methods:{
        LoginOut(){
      window.sessionStorage.clear()
      this.$router.push('/login')
        },
        selectAllStudent(row){
            console.log(row)
            axios({
                method:'get',
                url:globalurl.BASEURL + "StudentSelectAll",
            }).then(res=>{
                this.Stabledata = res.data.data
                this.studentVisible=true
            })
        },
        Openfrom(row){
            this.students.ClassName = row.ClassName
            this.students.Sname = row.Sname
            this.students.Sno=row.Sno
            this.changestudentVisible = true

        },
        selectAllClass(){
            axios({
                method:'get',
                url:globalurl.BASEURL + "ClassSelectAll",
            }).then(res=>{
                this.Classtabledata = res.data.data
                this.classVisible=true
            })
        },
        selectAllTeacher(){
            axios({
                method:'get',
                url:globalurl.BASEURL + "TeacherSelectAll",
            }).then(res=>{
                this.Ttabledata = res.data.data
                this.teacherVisible=true
            })
            
        },
        selectAllCourses(){
            axios({
                method:'get',
                url:globalurl.BASEURL + "CourseSelectAll",
            }).then(res=>{
                this.Ctabledata = res.data.data
                this.courseVisible=true
            })
        },
        AddCourse(){
            axios({
          method:'post',
          url: globalurl.BASEURL + "CourseAdd",
          data:this.courses,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            alert("添加成功")
            //this.tableData.Status = true;
            this.classVisible=false;
          }else{
            alert("添加失败")
            this.classVisible=false;
          }
        })          
        },
        AddTeacher(){
            axios({
          method:'post',
          url: globalurl.BASEURL + "TeacherAdd",
          data:this.teachers,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            alert("上传成功")
            //this.tableData.Status = true;
            this.submitVisible=false;
          }else{
            alert("上传失败")
            this.submitVisible=false;
          }
        })          
        },
        AddClass(){
            axios({
          method:'post',
          url: globalurl.BASEURL + "ClassAdd",
          //更新地址
          data:this.Class,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            alert("上传成功")
            //this.tableData.Status = true;
            this.submitVisible=false;
          }else{
            alert("上传失败")
            this.submitVisible=false;
          }
        })          
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
        submitImportForm() {           //导入教师信息
            console.log("sac杀杀杀")
            const params = new FormData()
            this.fileList.forEach((x) => {
            params.append('file', x.file)
            });
            //params.append('Cno',this.CourseMsg.Cno)
            //这里根据自己封装的axios来进行调用后端接口
            axios({
                method:'post',
                //url:'http://localhost:9999/upload',
                url:'',
                //更新地址
                data:params,
                Headers:{"Content-Type":"multipart/FormData" },
                }).then(res=>res.data).then(res=>{
                    if (res.succession ==="true"){
                    alert("上传成功")
                    //this.tableData.Status = true;
                    this.submitVisible=false;
                }else{
                    alert("上传失败")
                    this.submitVisible=false;
                }
            })          
       },
       Delteacher(row){
        axios({
          method:'post',
          url: globalurl.BASEURL + "TeacherDelete",
          //更新地址
          data:row,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            alert("删除成功")
            //this.tableData.Status = true;
            this.teacherVisible=false;
          }else{
            alert("删除失败")
            this.teacherVisible=false;
          }
        })      
       },
       Delstudent(row){
        axios({
          method:'post',
          url: globalurl.BASEURL + "StudentDelete",
          //更新地址
          data:row,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            alert("删除成功")
            //this.tableData.Status = true;
            this.studentVisible=false;
          }else{
            alert("删除失败")
            this.studentVisible=false;
          }
        })      
       },
       Delcourse(row){
        axios({
          method:'post',
          url: globalurl.BASEURL + "CourseDelete",
          //更新地址
          data:row,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            alert("删除成功")
            //this.tableData.Status = true;
            this.courseVisible=false;
          }else{
            alert("删除失败")
            this.courseVisible=false;
          }
        })      
       },
       Delclass(row){
        console.log("ascvvv")
        axios({
          method:'post',
          url: globalurl.BASEURL + "ClassDelete",
          //更新地址
          data:row,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            alert("删除成功")
            //this.tableData.Status = true;
            this.classVisible=false;
          }else{
            alert("删除失败")
            this.classVisible=false;
          }
        }) 
       },
       ChangeStudent(){
        axios({
            method:'post',
            url: globalurl.BASEURL + "StudentUpdate",
          //更新地址
          data:this.students,
          Headers:{"Content-Type":"multipart/FormData" },
        }).then(res=>res.data).then(res=>{
          if (res.succession ==="true"){
            alert("修改成功")
            //this.tableData.Status = true;
            this.changestudentVisible=false;
          }else{
            alert("修改失败")
            this.changestudentVisible=false;
          }
        })
       },
    }
  };
</script>

<style>
  .el-header {
    background-color: #B3C0D1;
    color: #333;
    line-height: 60px;
    background-image: 'https://c-ssl.dtstatic.com/uploads/item/201811/06/20181106101737_upvlq.thumb.1000_0.jpg';
  }
  .el-container{

  }
  .el-aside {
    color: #333;
  }
</style>