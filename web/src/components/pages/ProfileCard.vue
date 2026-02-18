<template>
  <el-card class="profile-card">
    <div class="profile-header">
      <div class="avatar-wrapper">
        <el-image
          :src="websiteStore.state.websiteInfo.logo || '/image/avatar.jpg'"
          alt="avatar"
          fit="cover"
        >
          <template #error>
            <div class="avatar-placeholder">
              <el-icon><User /></el-icon>
            </div>
          </template>
        </el-image>
      </div>
      <h3 class="profile-name">{{ websiteStore.state.websiteInfo.name }}</h3>
      <p class="profile-job">{{ websiteStore.state.websiteInfo.job }}</p>
    </div>
    
    <div class="profile-info">
      <div class="info-item">
        <el-icon><Location /></el-icon>
        <span>{{ websiteStore.state.websiteInfo.address }}</span>
      </div>
      <div class="info-item">
        <el-icon><Message /></el-icon>
        <span>{{ websiteStore.state.websiteInfo.email }}</span>
      </div>
    </div>
    
    <div class="profile-description">
      <p>{{ websiteStore.state.websiteInfo.description }}</p>
    </div>
    
    <div class="contact-section" v-if="websiteStore.state.websiteInfo.qq_image || websiteStore.state.websiteInfo.wechat_image" style="display: none;">
      <div class="contact-item" v-if="websiteStore.state.websiteInfo.qq_image">
        <el-image :src="websiteStore.state.websiteInfo.qq_image" alt="QQ"/>
        <span>QQ</span>
      </div>
      <div class="contact-item" v-if="websiteStore.state.websiteInfo.wechat_image">
        <el-image :src="websiteStore.state.websiteInfo.wechat_image" alt="微信"/>
        <span>微信</span>
      </div>
    </div>
    
    <div class="social-links" style="display: none;">
      <el-tooltip content="GitHub" placement="top" v-if="websiteStore.state.websiteInfo.github_url">
        <a :href="websiteStore.state.websiteInfo.github_url" target="_blank" class="social-link">
          <el-icon><component is="Link" /></el-icon>
        </a>
      </el-tooltip>
      <el-tooltip content="Gitee" placement="top" v-if="websiteStore.state.websiteInfo.gitee_url">
        <a :href="websiteStore.state.websiteInfo.gitee_url" target="_blank" class="social-link">
          <el-icon><component is="Link" /></el-icon>
        </a>
      </el-tooltip>
      <el-tooltip content="Bilibili" placement="top" v-if="websiteStore.state.websiteInfo.bilibili_url">
        <a :href="websiteStore.state.websiteInfo.bilibili_url" target="_blank" class="social-link">
          <el-icon><VideoPlay /></el-icon>
        </a>
      </el-tooltip>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import {useWebsiteStore} from "@/stores/website";
import {User, Location, Message, VideoPlay} from '@element-plus/icons-vue';

const websiteStore = useWebsiteStore()
</script>

<style scoped lang="scss">
.profile-card {
  background: var(--bg-card);
  border-radius: var(--radius-lg);
  border: none;
  box-shadow: var(--shadow-sm);
  overflow: hidden;
  
  :deep(.el-card__body) {
    padding: 0;
  }
  
  .profile-header {
    text-align: center;
    padding: 30px 20px;
    background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-primary-light) 100%);
    
    .avatar-wrapper {
      width: 90px;
      height: 90px;
      margin: 0 auto 16px;
      border-radius: 50%;
      overflow: hidden;
      border: 3px solid rgba(255, 255, 255, 0.3);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
      
      .el-image {
        width: 100%;
        height: 100%;
      }
      
      .avatar-placeholder {
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        background: rgba(255, 255, 255, 0.2);
        color: #fff;
        font-size: 2rem;
      }
    }
    
    .profile-name {
      color: #fff;
      font-size: 1.25rem;
      font-weight: 600;
      margin: 0 0 4px 0;
    }
    
    .profile-job {
      color: rgba(255, 255, 255, 0.85);
      font-size: 0.9rem;
      margin: 0;
    }
  }
  
  .profile-info {
    padding: 20px;
    border-bottom: 1px solid var(--border-color-light);
    
    .info-item {
      display: flex;
      align-items: center;
      gap: 10px;
      padding: 8px 0;
      color: var(--text-secondary);
      font-size: 0.9rem;
      
      .el-icon {
        color: var(--color-primary);
        font-size: 1.1rem;
      }
    }
  }
  
  .profile-description {
    padding: 16px 20px;
    border-bottom: 1px solid var(--border-color-light);
    
    p {
      margin: 0;
      color: var(--text-secondary);
      font-size: 0.9rem;
      line-height: 1.6;
    }
  }
  
  .contact-section {
    display: flex;
    justify-content: center;
    gap: 30px;
    padding: 20px;
    border-bottom: 1px solid var(--border-color-light);
    
    .contact-item {
      text-align: center;
      
      .el-image {
        width: 80px;
        height: 80px;
        border-radius: var(--radius-sm);
        margin-bottom: 8px;
      }
      
      span {
        display: block;
        font-size: 0.85rem;
        color: var(--text-muted);
      }
    }
  }
  
  .social-links {
    display: flex;
    justify-content: center;
    gap: 16px;
    padding: 20px;
    
    .social-link {
      width: 40px;
      height: 40px;
      display: flex;
      align-items: center;
      justify-content: center;
      border-radius: 50%;
      background: var(--bg-secondary);
      color: var(--text-secondary);
      transition: all 0.3s ease;
      
      &:hover {
        background: var(--color-primary);
        color: #fff;
        transform: translateY(-2px);
      }
      
      .el-icon {
        font-size: 1.2rem;
      }
    }
  }
}
</style>
