<template>
    <div style="background-color: #dededf">
        <my-title :active-index="'2'"></my-title>
        <div style="min-height: 650px">
            <student-list style="width: 48%;margin-left: 2%;margin-top: 20px;float: left"></student-list>
            <add style="width: 48%;margin-left: 1%;margin-top: 20px;float: left"></add>
        </div>
    </div>
</template>

<script>
    import MyTitle from "../../components/myTitle";
    import StudentList from "../../components/studentList";
    import Add from "../../components/add";
    export default {
        name: "addStudent",
        components: {Add, StudentList, MyTitle}
    }
</script>

<style scoped>

</style>