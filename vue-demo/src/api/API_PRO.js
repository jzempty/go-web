export default {

    //baseURL: 'http://39.106.160.119:8080/',//后端地址
    baseURL: 'http://127.0.0.1:8080/',//后端地址
    
    method: 'post',
    //登录模块接口
    studentLogin: {url:'student/login'},
    teacherLogin:{url:'teacher/login'},

    //注册模块
    studentReg:{url:'student/insertStudent'},
    teacherReg:{url:'teacher/insertTeacher'},

    getStuList:{url:'student/getStudentList'},

    allHomeworkList:{url:'homework/getHomeworkList'},
    tHomeworkList:{url:'homework/getTHomeworkList'},

    addHW:{url:'homework/addHomework'},
    deletHW:{url:'homework/deleteHomework'},
    updateHW:{url:'homework/updateHome'},

    detailSHList:{url:'homework/getDetailHomeworkList'},
    studentSHList:{url:'homework/getMyHomeworkList'},

    submitSHom:{url:'homework/insertHomework'},
    updateSHome:{url:'homework/updateHomework'},


}
