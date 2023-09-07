<template>
    <div class="report">
      <h2>实验报告</h2>
      <el-table
        :data="reports"
        style="width: 100%">
        <el-table-column prop="name" label="姓名"></el-table-column>

        <el-table-column prop="reportUrl" label="报告">
          <template slot-scope="scope">
            <el-link :href="scope.row.reportUrl"  target="_blank">下载</el-link>
          </template>
        </el-table-column>
        <el-table-column
          label="操作">
          <template slot-scope="scope">
            <el-button
              type="primary"
              @click="downloadReport(scope.row)">
              下载报告
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
</template>
  

  <script>
  export default {
    data() {
      return {
        reports: [
          {
            name: '张三',
            reportUrl: 'http://www.example.com/report1.pdf'
          },
          {
            name: '李四',
            reportUrl: 'http://www.example.com/report2.pdf'
          },
          {
            name: '王五',
            reportUrl: 'http://www.example.com/report3.pdf'
          }
        ]
      }
    },
    methods: {
      downloadReport(report) {
        let xhr = new XMLHttpRequest();
        xhr.open('GET', report.reportUrl, true);
        xhr.responseType = 'blob';
  
        xhr.onload = () => {
          if (xhr.status === 200) {
            let blob = new Blob([xhr.response], { type: 'application/pdf' });
            let url = window.URL.createObjectURL(blob);
  
            let link = document.createElement('a');
            link.href = url;
            link.download = `${report.name}实验报告.pdf`;
            link.click();
  
            window.URL.revokeObjectURL(url);
          }
        };
  
        xhr.send();
      }
    }
  }
  </script>
  