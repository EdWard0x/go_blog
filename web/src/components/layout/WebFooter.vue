<template>
  <div class="web-footer">
    <div class="container">
      <div class="footer-main">
        <div class="footer-brand">
          <div class="logo">
            <el-image
                :src="websiteStore.state.websiteInfo.full_logo===''?'/image/full_logo.png':websiteStore.state.websiteInfo.full_logo"
                alt=""/>
          </div>
          <p class="description">{{ websiteStore.state.websiteInfo.description }}</p>
        </div>
        
        <div class="footer-links">
          <div class="link-group" style="display:none">
            <h4>快速链接</h4>
            <div class="links">
              <el-link v-for="item in footerLinkList" :key="item.title" :href="item.link">
                {{ item.title }}
              </el-link>
            </div>
          </div>
          
          <div class="link-group">
            <h4>社交媒体</h4>
            <div class="social-links">
              <el-tooltip content="GitHub" placement="top" v-if="websiteStore.state.websiteInfo.github_url">
                <a :href="websiteStore.state.websiteInfo.github_url" target="_blank" class="social-link">
                  <el-image src="/image/github.png" alt="GitHub"/>
                </a>
              </el-tooltip>
              <el-tooltip content="Gitee" placement="top" v-if="websiteStore.state.websiteInfo.gitee_url">
                <a :href="websiteStore.state.websiteInfo.gitee_url" target="_blank" class="social-link">
                  <el-image src="/image/gitee.png" alt="Gitee"/>
                </a>
              </el-tooltip>
              <el-tooltip content="Bilibili" placement="top" v-if="websiteStore.state.websiteInfo.bilibili_url">
                <a :href="websiteStore.state.websiteInfo.bilibili_url" target="_blank" class="social-link">
                  <el-image src="/image/bilibili.png" alt="Bilibili"/>
                </a>
              </el-tooltip>
            </div>
          </div>
          
          <div class="link-group" style="display:none">
            <h4>备案信息</h4>
            <div class="filing-info">
              <el-image src="/image/filing.png" alt=""/>
              <el-link href="https://beian.miit.gov.cn/#/Integrated/index" :underline="false">
                {{ websiteStore.state.websiteInfo.icp_filing }}
              </el-link>
              <el-link :href="publicSecurityFilingLink" :underline="false">
                {{ websiteStore.state.websiteInfo.public_security_filing }}
              </el-link>
            </div>
          </div>
        </div>
        
        <div class="footer-info">
          <div class="info-item">
            <el-icon><Calendar /></el-icon>
            <span>建站日期：{{ websiteStore.state.websiteInfo.created_at }}</span>
          </div>
          <div class="info-item run-time">
            <el-icon><Timer /></el-icon>
            <span>已运行：{{ elapsedTime }}</span>
          </div>
          <div class="version-info">
            <el-tag type="primary" size="small" effect="plain">v{{ websiteStore.state.websiteInfo.version }}</el-tag>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useWebsiteStore} from "@/stores/website";
import {computed} from "vue";
import {ref} from "vue";
import {onUnmounted} from "vue";
import {type FooterLink, websiteFooterLink} from "@/api/website";
import {Calendar, Timer} from '@element-plus/icons-vue';

const footerLinkList=ref<FooterLink[]>([])

const getFooterLinkList =async ()=>{
  const res = await websiteFooterLink()
  if (res.code===0){
    footerLinkList.value=res.data
  }
}

getFooterLinkList()

const websiteStore = useWebsiteStore()

let timerId: number | null = null;
const elapsedTime = ref("");

function updateElapsedTime() {
  let creationDate = websiteStore.state.websiteInfo.created_at;
  if (creationDate) {
    let creationTimestamp = new Date(creationDate).getTime();
    let currentTimestamp = new Date().getTime();

    let totalDays = (currentTimestamp - creationTimestamp) / 1000 / (60 * 60 * 24);
    let daysPassed = Math.floor(totalDays);
    let hoursRemaining = Math.floor((totalDays - daysPassed) * 24);
    let minutesRemaining = Math.floor((totalDays - daysPassed - (hoursRemaining / 24)) * 24 * 60);
    let secondsRemaining = Math.floor((totalDays - daysPassed - (hoursRemaining / 24) - (minutesRemaining / 24 / 60)) * 24 * 60 * 60);

    elapsedTime.value = `${daysPassed}天${hoursRemaining}时${minutesRemaining}分${secondsRemaining}秒`;
  }
}

function initializeTimer() {
  updateElapsedTime();
  timerId = setInterval(updateElapsedTime, 1000);
}

onUnmounted(() => {
  clearInterval(timerId as number);
});

initializeTimer();

const publicSecurityFilingLink = computed(() => "http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=" + websiteStore.state.websiteInfo.public_security_filing.match(/\d+/))
</script>


<style scoped lang="scss">
.web-footer {
  background: linear-gradient(180deg, var(--bg-primary) 0%, var(--bg-secondary) 100%);
  border-top: 1px solid var(--border-color-light);
  padding: 40px 20px 20px;

  .container {
    max-width: 1400px;
    margin: 0 auto;
    
    .footer-main {
      display: grid;
      grid-template-columns: 1.5fr 1.5fr 1fr;
      gap: 40px;
      
      .footer-brand {
        .logo {
          margin-bottom: 16px;
          
          .el-image {
            height: 50px;
          }
        }
        
        .description {
          color: var(--text-secondary);
          font-size: 0.9rem;
          line-height: 1.6;
          margin: 0;
        }
      }
      
      .footer-links {
        display: flex;
        flex-direction: column;
        gap: 24px;
        
        .link-group {
          h4 {
            color: var(--text-primary);
            font-size: 1rem;
            font-weight: 600;
            margin: 0 0 12px 0;
          }
          
          .links {
            display: flex;
            flex-direction: column;
            gap: 8px;
            
            .el-link {
              color: var(--text-secondary);
              font-size: 0.9rem;
              transition: color 0.2s ease;
              
              &:hover {
                color: var(--color-primary);
              }
            }
          }
          
          .social-links {
            display: flex;
            gap: 12px;
            
            .social-link {
              display: flex;
              align-items: center;
              justify-content: center;
              width: 36px;
              height: 36px;
              border-radius: 50%;
              background: var(--bg-card);
              transition: all 0.3s ease;
              
              &:hover {
                background: var(--color-primary);
                transform: translateY(-2px);
              }
              
              .el-image {
                width: 20px;
                height: 20px;
              }
            }
          }
          
          .filing-info {
            display: flex;
            align-items: center;
            gap: 12px;
            flex-wrap: wrap;
            
            .el-image {
              height: 20px;
            }
            
            .el-link {
              color: var(--text-muted);
              font-size: 0.85rem;
              
              &:hover {
                color: var(--color-primary);
              }
            }
          }
        }
      }
      
      .footer-info {
        .info-item {
          display: flex;
          align-items: center;
          gap: 8px;
          color: var(--text-secondary);
          font-size: 0.9rem;
          margin-bottom: 12px;
          
          .el-icon {
            color: var(--color-primary);
          }
          
          &.run-time {
            font-family: 'Courier New', monospace;
          }
        }
        
        .version-info {
          margin-top: 16px;
          
          .el-tag {
            border-radius: 20px;
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .web-footer {
    .container {
      .footer-main {
        grid-template-columns: 1fr;
        gap: 24px;
      }
    }
  }
}
</style>
