<template>
  <div>
    <h1>课程分页</h1>
    <button @click="sub">分页按钮</button>
    <button @click="created">数据按钮</button>
    </div>
</template>
<script>
  
  import axios from 'axios'
  export default {
    methods: {
      sub() {
        
      },
      created() {
        //第二种写法
        axios({
          method: 'get',
          url: 'http://127.0.0.1:8080/login.html',
          params: {
            id: 1
          }
        }).then((res) => {
          this.test=res.data.data
          this.arr=res.data.a
          console.log("data",JSON.stringify(res.data))
          console.log("data[0]",JSON.stringify(res.data.data[0]))
          for (let i = 0; i < this.test.length; i++) {
            console.log("sno=",this.test[i].st);
            console.log("arr=",this.arr[i]);
          }
          alert(res.data.data)

          // console.log(arr);
        });
      }
    },

    data(){
      return{
        test:[{
          st:"",
          sno:"",
        }],
        arr:[]
      }
    },
  };

</script>
<style>
</style>