<template>
    <div style="background-color: #dededf">
        <my-title :active-index="'1'"></my-title>
        <div style="min-height: 650px">
            <el-card style="margin-left: 2%;width: 47%;margin-top: 20px;float: left">
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
                        <div class="tableTitle">
                            作业内容
                        </div>
                        <div class="tableTitle">
                            截止时间
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
                            {{item.hddl}}
                        </div>
                        <div class="tableCell">
                            <el-button v-if="item.isEnd===false && item.isStart===true"
                                       type="primary" size="mini" icon="el-icon-edit" circle @click="submitHomework(item)"></el-button>

                            <p  v-if="item.isStart===false">作业暂未开始</p>
                            <p v-if="item.isEnd===true">作业提交结束</p>
                        </div>
                    </div>

                </div>
            </el-card>

            <el-card style="margin-left: 2%;width: 47%;margin-top: 20px;float: left">
                <div style="display: table;width: 100%">
                    <div style="display: table-row">
                        <div class="tableTitle">
                            序号
                        </div>
                        <div class="tableTitle">
                            作业编号
                        </div>
                        <div class="tableTitle">
                            作业内容
                        </div>
                        <div class="tableTitle">
                        </div>
                    </div>
                    <div
                            :class="{'grayLine': (index % 2 === 0)}"
                            style="display: table-row;"
                            v-for="(item, index) in myHomeworkList"
                            :key="item.id">
                        <div class="tableCell">
                            {{index+1}}
                        </div>
                        <div class="tableCell">
                            {{item.homeworkId}}
                        </div>
                        <div class="tableCell">
                            {{item.homeworkContent}}
                        </div>
                        <div class="tableCell">
                            <el-button v-if="item.isEnd===false && item.isStart===true" type="primary" size="mini" @click="updateHomework(item)">修改</el-button>
                            <p v-if="item.isEnd===true">作业提交结束</p>
                        </div>
                    </div>
                </div>
            </el-card>

            <el-dialog title="提交作业" :visible.sync="submitVisible">
                <el-form :model="form1">
                    <el-form-item label="作业序号" :label-width="'120px'">
                        <el-input v-model="form.hid" autocomplete="off" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="作业名称" :label-width="'120px'">
                        <el-input v-model="form.name" autocomplete="off" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="作业内容" :label-width="'120px'">
                        <el-input v-model="form.content" autocomplete="off" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="填写作业" :label-width="'120px'">
                        <el-input
                                type="textarea"
                                :rows="5"
                                placeholder="请输入内容"
                                v-model="form.homework_content"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="submitVisible = false">取 消</el-button>
                    <el-button type="primary" @click="submit">确 定</el-button>
                </div>
            </el-dialog>

            <el-dialog title="更新作业" :visible.sync="updateVisible">
                <el-form :model="form1">
                    <el-form-item label="作业序号" :label-width="'120px'">
                        <el-input v-model="form1.hid" autocomplete="off" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="修改作业" :label-width="'120px'">
                        <el-input
                                type="textarea"
                                :rows="5"
                                placeholder="请输入内容"
                                v-model="form1.homework_content"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="updateVisible = false">取 消</el-button>
                    <el-button type="primary" @click="update">确 定</el-button>
                </div>
            </el-dialog>

            <my-footer style="bottom: 0;position: absolute;z-index: 999;"></my-footer>
        </div>
    </div>
</template>

<script>
    import API from "../../api";
    import Cookies from 'js-cookie'
    import qs from "qs";
    import MyTitle from "../../components/myTitle";
    import MyFooter from "../../components/myFooter";
    export default {
        name: "index",
        components: {MyFooter, MyTitle},
        data(){
            return{
                sid:Cookies.get('sid'),
                homeworkList:[],
                myHomeworkList:[],
                submitVisible:false,
                updateVisible:false,
                certainId:0,
                form:{
                    hid:0,
                    name:'',
                    content:'',
                    homework_content:"",
                },
                form1:{
                    hid:0,
                    homework_content:"",
                }
            }
        },

        mounted(){
            this.getHomeworkList();
            this.getMyHomeworkList();
        },

        methods:{
            getHomeworkList(){
                let data={
                };
                API.allHomeworkList(data).then(res=>{
                    this.homeworkList=res;
                    let dd = new Date();

                    if(this.homeworkList.length!==0){
                        this.homeworkList.forEach(item=>{
                            let start=new Date(item.startTime*1000);
                            let end=new Date(item.endTime*1000);
                            item.hddl=new Date(item.endTime*1000).toLocaleDateString();
                            if(start<dd && end>dd){
                                item.isStart=true;
                                item.isEnd=false;
                            }else if(start>dd){
                                item.isStart=false;
                                item.isEnd=false;
                            }else{
                                item.isStart=true;
                                item.isEnd=true;
                            }
                        })
                    }
                }).catch(msg=>{
                    alert(msg);
                });
            },

            submitHomework(item){
                this.form.hid=item.hid;
                this.form.name=item.htitle;
                this.form.content=item.hcontent;
                this.submitVisible=true;
            },

            submit(){
                let data = {
                    sid:parseInt(this.sid),
                    hid:parseInt(this.form.hid),
                    homework_content:this.form.homework_content.toString(),
                    create_time: Math.round(new Date().getTime()/1000) ,
                };
                data = qs.stringify(data);
                API.submitSHom(data).then(res=>{
                    if(res==='Successful Insert'){
                        alert("提交作业成功");
                        this.form.homework_content="";
                        this.getMyHomeworkList();
                    }else{
                        alert("提交作业失败");
                    }
                    this.submitVisible=false;
                }).catch(msg=>{
                    alert(msg);
                })
            },

            getMyHomeworkList(){
                let data={
                    sid:parseInt(this.sid),
                };
                data = qs.stringify(data);
                API.studentSHList(data).then(res=>{
                    this.myHomeworkList=res;
                    if(this.myHomeworkList.length!==0){
                        this.myHomeworkList.forEach(item=>{
                            for(let i=0;i<this.homeworkList.length;i++){
                                if(item.homeworkId===this.homeworkList[i].hid){
                                    item.isEnd=this.homeworkList[i].isEnd;
                                    item.isStart=this.homeworkList[i].isStart;
                                }
                            }
                        })
                    }
                }).catch(msg=>{
                    alert(msg);
                });
            },

            updateHomework(item){
                this.form1.hid=item.homeworkId;
                this.form1.homework_content=item.homeworkContent;
                this.certainId=item.id;
                this.updateVisible=true;
            },

            update(){
                let data = {
                    id:parseInt(this.certainId),
                    homework_content:this.form1.homework_content,
                    update_time: Math.round(new Date().getTime()/1000) ,
                };
                data = qs.stringify(data);
                API.updateSHome(data).then(res=>{
                    if(res==='Successful update'){
                        alert("修改作业成功");
                        this.getMyHomeworkList();
                    }else{
                        alert("修改作业失败");
                    }
                    this.updateVisible=false;
                }).catch(msg=>{
                    alert(msg);
                })
            },


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