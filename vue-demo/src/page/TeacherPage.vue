<template>
    <div style="background-color: #dededf">
        <my-title :active-index="'1'"></my-title>
        <div style="min-height: 650px">
            <el-card style="margin: auto;width: 1000px;margin-top: 20px">
                <el-button type="success" style="float: right;margin-right: 20px;margin-bottom: 10px" @click="newHomeworkVisible=true">添加作业</el-button>
                <div style="display: table;width: 100%">
                    <div style="display: table-row">
                        <div class="tableTitle">
                            序号
                        </div>
                        <div class="tableTitle">
                            作业编号
                        </div>
                        <div class="tableTitle">
                            作业名称
                        </div>
                        <div class="tableTitle" style="width: 300px">
                            作业内容
                        </div>
                        <div class="tableTitle">
                            开始时间
                        </div>
                        <div class="tableTitle">
                            截止时间
                        </div>
                        <div class="tableTitle">
                        </div>
                        <div class="tableTitle">
                        </div>
                        <div class="tableTitle">
                        </div>
                    </div>
                    <div
                            :class="{'grayLine': (index % 2 === 0)}"
                            style="display: table-row;"
                            v-for="(item, index) in homeworkList"
                            :key="item.id">
                        <div class="tableCell">
                            {{index+1}}
                        </div>
                        <div class="tableCell">
                            {{item.hid}}
                        </div>
                        <div class="tableCell">
                            {{item.htitle}}
                        </div>
                        <div class="tableCell">
                            {{item.hcontent}}
                        </div>
                        <div class="tableCell">
                            {{item.hstart}}
                        </div>
                        <div class="tableCell">
                            {{item.hddl}}
                        </div>
                        <div class="tableCell">
                            <el-button size="mini" type="primary" @click="updateHomework(item)">修改</el-button>
                        </div>
                        <div class="tableCell">
                            <el-button size="mini" type="primary" @click="delH(item.hid)">删除</el-button>
                        </div>
                        <div class="tableCell">
                            <el-button size="mini" type="primary" @click="checkHomeworkList(item.hid,item.htitle)">查看提交</el-button>
                        </div>
                    </div>

                </div>
            </el-card>

            <el-dialog
                    title="提示"
                    :visible.sync="deleteVisible"
                    width="30%">
                <span>确定删除该作业吗</span>
                <span slot="footer" class="dialog-footer">
                    <el-button @click="deleteVisible = false">取 消</el-button>
                    <el-button type="primary" @click="deleteHomework">确 定</el-button>
                </span>
            </el-dialog>

            <el-dialog :title="nowHomeworkTitle" :visible.sync="showHomeworkList">
                <div style="display: table;width: 100%">
                    <div style="display: table-row">
                        <div class="tableTitle">
                            序号
                        </div>
                        <div class="tableTitle">
                            学生学号
                        </div>
                        <div class="tableTitle" style="width: 200px">
                            提交内容
                        </div>
                        <div class="tableTitle">
                            创建时间
                        </div>
                        <div class="tableTitle">
                            更新时间
                        </div>
                    </div>
                    <div
                            :class="{'grayLine': (index % 2 === 0)}"
                            style="display: table-row;"
                            v-for="(item, index) in studentHomeworkList"
                            :key="item.id">
                        <div class="tableCell">
                            {{index+1}}
                        </div>
                        <div class="tableCell">
                            {{item.studentId}}
                        </div>
                        <div class="tableCell">
                            {{item.homeworkContent}}
                        </div>
                        <div class="tableCell">
                            {{item.time1}}
                        </div>
                        <div class="tableCell">
                            {{item.time2}}
                        </div>
                    </div>

                </div>
                <span slot="footer" class="dialog-footer">
                    <el-button @click="showHomeworkList = false">取 消</el-button>
                    <el-button type="primary" @click="showHomeworkList = false">确 定</el-button>
                </span>
            </el-dialog>

            <el-dialog title="添加作业" :visible.sync="newHomeworkVisible">
                <el-form :model="newHomework">
                    <el-form-item label="作业标题" :label-width="'120px'">
                        <el-input v-model="newHomework.title" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="作业内容" :label-width="'120px'">
                        <el-input v-model="newHomework.content" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="时间范围" :label-width="'120px'">
                        <el-date-picker
                                v-model="newHomework.time"
                                type="daterange"
                                range-separator="至"
                                start-placeholder="开始日期"
                                end-placeholder="结束日期">
                        </el-date-picker>
                    </el-form-item>

                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="newHomeworkVisible = false">取 消</el-button>
                    <el-button type="primary" @click="addHomework">确 定</el-button>
                </div>
            </el-dialog>

            <el-dialog title="修改作业" :visible.sync="updateHomeworkVisible">
                <el-form :model="newHomework">
                    <el-form-item label="作业标题" :label-width="'120px'">
                        <el-input v-model="newHomework.title" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="作业内容" :label-width="'120px'">
                        <el-input v-model="newHomework.content" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="时间范围" :label-width="'120px'">
                        <el-date-picker
                                v-model="newHomework.time"
                                type="daterange"
                                range-separator="至"
                                start-placeholder="开始日期"
                                end-placeholder="结束日期">
                        </el-date-picker>
                    </el-form-item>

                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="updateHomeworkVisible = false">取 消</el-button>
                    <el-button type="primary" @click="update">确 定</el-button>
                </div>
            </el-dialog>

        </div>
        <my-footer style="bottom: 0;position: absolute"></my-footer>
    </div>
</template>

<script>
    //import MyTitle from "../components/myTitle";
    //import MyFooter from "../components/myFooter";
    import Cookies from 'js-cookie'
    import API from "../api";
    import qs from "qs";
    export default {
        name: "TeacherPage",
        //components: {MyFooter, MyTitle},
        data(){
            return{
                tid:Cookies.get('tid'),
                homeworkList:[],
                newHomework:{
                    title:'a',
                    content:'b',
                    time:'2022',
                },
                showHomeworkList:false,
                nowHomeworkTitle:"",
                noData:false,
                newHomeworkVisible:false,
                deleteVisible:false,
                delID:0,
                updateId:0,
                studentHomeworkList:[],
                updateHomeworkVisible:false,
            }
        },

        mounted(){
            this.getHomeworkList();
        },

        methods:{
            checkHomeworkList(id, title){
                let data = {
                    hid:id,
                };
                data = qs.stringify(data);

                API.detailSHList(data).then(res=>{
                    this.studentHomeworkList=res;
                    if(this.studentHomeworkList.length!==0){
                        this.studentHomeworkList.forEach(item=>{
                            item.time1= new Date(item.createTime * 1000).toLocaleDateString();
                            item.time2= new Date(item.updateTime * 1000).toLocaleDateString();
                        })
                    }
                }).catch(msg=>{
                    alert(msg);
                });
                this.nowHomeworkTitle=title;
                this.showHomeworkList=true;
            },

            getHomeworkList(){
                let data = {
                    tid:parseInt(this.tid),
                };
                data = qs.stringify(data);
                API.tHomeworkList(data).then(res=>{
                    this.homeworkList=res;
                    if(this.homeworkList.length!==0){
                        this.homeworkList.forEach(item=>{
                            item.hstart= new Date(item.startTime * 1000).toLocaleDateString();
                            item.hddl= new Date(item.endTime * 1000).toLocaleDateString();
                        })
                    }
                }).catch(msg=>{
                    alert(msg);
                })
            },

            addHomework(){
                let data = {
                    tid:parseInt(this.tid),
                    htitle:this.newHomework.title,
                    hcontent:this.newHomework.content,
                    start_time: new Date(this.newHomework.time[0]).valueOf() / 1000,
                    end_time: new Date(this.newHomework.time[1]).valueOf() / 1000,
                };
                data = qs.stringify(data);
                API.addHW(data).then(res=>{
                    if(res==='Successful insert'){
                        alert("添加作业成功");
                        this.getHomeworkList();
                        this.newHomework.htitle="";
                        this.newHomework.hcontent="";
                        this.newHomework.time="";
                    }else{
                        alert("添加作业失败");
                    }
                    this.newHomeworkVisible=false;
                }).catch(msg=>{
                    alert(msg);
                })
            },

            delH(hid){
                this.deleteVisible=true;
                this.delID=hid;
            },

            deleteHomework(){
                let data={
                    hid:this.delID,
                };
                data = qs.stringify(data);
                API.deletHW(data).then(res=>{
                    if(res==='Successful delete'){
                        alert("删除作业成功");
                        this.getHomeworkList();
                    }else{
                        alert("删除作业失败");
                    }
                    this.deleteVisible=false;
                    this.delID=0;
                }).catch(msg=>{
                    alert(msg);
                })
            },

            updateHomework(item){
                this.newHomework.title=item.htitle;
                this.newHomework.content=item.hcontent;
                this.updateId=item.hid;
                this.updateHomeworkVisible=true;
            },

            update(){
                let data = {
                    hid:this.updateId,
                    htitle:this.newHomework.title,
                    hcontent:this.newHomework.content,
                    start_time: new Date(this.newHomework.time[0]).valueOf() / 1000,
                    end_time: new Date(this.newHomework.time[1]).valueOf() / 1000,
                };
                data = qs.stringify(data);
                API.updateHW(data).then(res=>{
                    if(res==='Successful update'){
                        alert("修改作业内容成功");
                        this.getHomeworkList();
                        this.newHomework.htitle="";
                        this.newHomework.hcontent="";
                        this.newHomework.time="";
                    }else{
                        alert("修改作业失败");
                    }
                    this.updateHomeworkVisible=false;
                }).catch(msg=>{
                    alert(msg);
                })
            }
        }
    }
</script>

<style scoped>
    .tableTitle{
        display: table-cell;
        text-align: center;
        font-size: 15px;
        padding: 5px;
        color: gray;
    }

    .tableCell{
        display: table-cell;
        text-align: center;
        font-size: 15px;
        padding: 20px 10px;
    }

    .grayLine{
        background-color: #f5f5f5;
    }
</style>