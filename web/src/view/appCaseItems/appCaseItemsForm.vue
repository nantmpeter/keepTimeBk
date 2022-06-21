<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="caseId字段:">
          <el-input v-model.number="formData.caseId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="userId字段:">
          <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="date字段:">
          <el-date-picker v-model="formData.date" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="start字段:">
          <el-date-picker v-model="formData.start" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="end字段:">
          <el-date-picker v-model="formData.end" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AppCaseItems'
}
</script>

<script setup>
import {
  createAppCaseItems,
  updateAppCaseItems,
  findAppCaseItems
} from '@/api/appCaseItems'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        caseId: 0,
        userId: 0,
        date: new Date(),
        start: new Date(),
        end: new Date(),
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAppCaseItems({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappCaseItems
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createAppCaseItems(formData.value)
          break
        case 'update':
          res = await updateAppCaseItems(formData.value)
          break
        default:
          res = await createAppCaseItems(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
