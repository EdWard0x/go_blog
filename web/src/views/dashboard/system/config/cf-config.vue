<template>
  <div class="cf-config">
    <el-col :span="12">
      <div class="info">
        <div class="title">
          <el-row>cloudflare配置</el-row>
        </div>
        <div class="content">
          <el-form
              :model="cfInfo"
              :validate-on-rule-change="false"
              label-width="auto"
              style="max-width: 400px"
          >
            <el-form-item label="访问密钥ID">
              <el-input @change="updateCfInfo" v-model="cfInfo.accessKeyId" type="password" show-password/>
            </el-form-item>
            <el-form-item label="访问密钥Key">
              <el-input @change="updateCfInfo" v-model.number="cfInfo.accessKeySecret" type="password" show-password/>
            </el-form-item>
            <el-form-item label="账户ID">
              <el-input @change="updateCfInfo" v-model="cfInfo.accountId" type="password" show-password/>
            </el-form-item>
            <el-form-item label="存储桶名称">
              <el-input @change="updateCfInfo" v-model="cfInfo.bucketName"/>
            </el-form-item>
            <el-form-item label="CDN域名">
              <el-input @change="updateCfInfo" v-model="cfInfo.cdn"/>
            </el-form-item>
<!--            <el-form-item label="使用CDN上传加速">-->
<!--              <el-switch v-model="cfInfo.use_cdn_domains" @change="updateQiniuInfo"/>-->
<!--            </el-form-item>-->
<!--            <el-form-item label="是否启用">-->
<!--              <el-switch v-model="cfInfo.enabled" @change="updateCfInfo"/>-->
<!--            </el-form-item>-->
          </el-form>
        </div>
      </div>
    </el-col>
  </div>
</template>

<script setup lang="ts">
import {ref, watch} from "vue";
// import {type Qiniu, getQiniu, updateQiniu} from "@/api/config";
import {type Cf, getCf, updateCf} from "@/api/config";

const cfInfo = ref<Cf>({
  accessKeyId: '',
  accessKeySecret: '',
  accountId: '',
  bucketName: '',
  cdn: '',
})

const getCfInfo = async () => {
  const res = await getCf()
  if (res.code === 0) {
    cfInfo.value = res.data
  }
}

getCfInfo()

const shouldRefreshInfo = ref(false)
watch(() => shouldRefreshInfo.value, (newVal) => {
  if (newVal) {
    getCfInfo()
    shouldRefreshInfo.value = false
  }
})

const updateCfInfo = async () => {
  const res = await updateCf(cfInfo.value)
  console.log(cfInfo.value)
  if (res.code === 0) {
    ElMessage.success(res.msg)
  } else {
    shouldRefreshInfo.value = true
  }
}
</script>

<style scoped lang="scss">
.cf-config {
  display: flex;

  .info {
    .title {
      border-left: 5px solid blue;
      padding-left: 10px;
    }

    .content {
      margin: 20px;
    }
  }
}
</style>
