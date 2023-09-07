//该文件专门创建整个应用的路由器
import VueRouter from 'vue-router'
 
import Vue from 'vue'
 
//引入组件
import About from '../views/AboutMe'
import Home from '../views/HomeMe'
import Login from '../views/LoginMe.vue'
import Register from '../views/RegisterMe'
import Teacher from '../views/TeacherPage'
import Student from '../views/StudentMe'
import Admin from '../views/AdminMe'
import Stu from '../views/StudentPage'
import Title from '../components/myTitle'
import TeacherPage from '../page/TeacherPage'
import Assign from '../views/AssignmentMe'
import StuAssign from '../views/StudentAssignment'
import DownLoad from '../page/downLoadMe'
import StuList from '../views/StudentsList'
import Test from '../components/TestMe'
import CPass from '../views/ChangePassword'
import BEmail from '../views/BindEmail'
import CP from '../temp/CP'

 
// 将组件挂载在实例上
Vue.use(VueRouter)

 
//创建并暴露一个路由器
const router=new VueRouter({
    routes:[
        {
            path:'/about',
            component:About
        },
        {
            path:'/cp',
            component:CP
        },
        {
            path:'/',
            component:Test,
        },
        {
            path:'/home',
            component:Home
        },
        {
            path:'/admin',
            component:Admin,
        },
        {
            path:'/register',
            component:Register,
        },
        {
            path:'/login',
            component:Login,
        },
        {
            path:'/changepassword',
            component:CPass,
        },
        {
            path:'/bindemail',
            component:BEmail,
        },
        {
            path:'/students',           //教师查看课程对应的学生信息
            component:StuList,
        },
        {
            path:'/student',
            component:Student,
        },
        {
            path:'/teacher',
            component:Teacher,
        },
        {
            path:'/tpage',
            component:TeacherPage
        },
        {
            path:'/down',
            component:DownLoad,
        },
        {
            path:'/assignment',
            component:Assign
        },
        {
            path:'/studentassign',
            component:StuAssign
        },
        {
            path:'/stu',
            component:Stu,
        },
        {
            path:'/title',
            component:Title,
        },
        
    ],
    mode:'history',
    
})
 
// 向外抛出路由组件
export default router
