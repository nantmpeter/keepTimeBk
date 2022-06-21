<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户UUID:">
          <el-input v-model="formData.uuid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户登录名:">
          <el-input v-model="formData.username" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户登录密码:">
          <el-input v-model="formData.password" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户昵称:">
          <el-input v-model="formData.nickName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户头像:">
          <el-input v-model="formData.headerImg" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户角色ID:">
          <el-input v-model="formData.authorityId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户手机号:">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户邮箱:">
          <el-input v-model="formData.email" clearable placeholder="请输入" />
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
  name: 'AppUsers'
}
</script>

<script setup>
import {
  createAppUsers,
  updateAppUsers,
  findAppUsers
} from '@/api/appUsers'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        uuid: '',
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: '',
        phone: '',
        email: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAppUsers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappUsers
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
          res = await createAppUsers(formData.value)
          break
        case 'update':
          res = await updateAppUsers(formData.value)
          break
        default:
          res = await createAppUsers(formData.value)
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
