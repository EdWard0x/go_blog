<template>
  <div :class="{'web-navbar': true, 'is-transparent': !isScrolled, 'is-solid': isScrolled}">
    <div class="container">
      <!-- <logo/> -->
      <div class="web-menu">
        <el-menu mode="horizontal" :ellipsis="false" :router="true" :default-active="$route.path">
          <template v-for="item in menuList">
            <el-menu-item :index="item.name"><span>{{ item.title }}</span></el-menu-item>
          </template>
        </el-menu>
      </div>
      <auth-popover/>
    </div>
  </div>
</template>

<script setup lang="ts">
import AuthPopover from "@/components/common/AuthPopover.vue";
import Logo from "@/components/widgets/Logo.vue";
import {ref} from "vue";
import {onMounted, onUnmounted} from "vue";

const isScrolled = ref(false)

const props = defineProps<{
  noScroll?: boolean
}>()

const handleScroll = () => {
  isScrolled.value = window.scrollY > 50
}

onMounted(() => {
  if (!props.noScroll) {
    window.addEventListener("scroll", handleScroll)
  }
})

onUnmounted(() => {
  if (!props.noScroll) {
    window.removeEventListener("scroll", handleScroll)
  }
})

interface MenuItem {
  title: string;
  name: string;
}

const menuList: MenuItem[] = [
  {
    title: "首页",
    name: "/",
  },
  {
    title: "搜索",
    name: "/search",
  },
  // {
  //   title: "新闻",
  //   name: "/news",
  // },
  // {
  //   title: "友链",
  //   name: "/friend-link",
  // },
  {
    title: "关于",
    name: "/about",
  }
]

</script>


<style scoped lang="scss">
.web-navbar {
  display: flex;
  justify-content: center;
  width: 100%;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 100;
  transition: all 0.3s ease;
  
  &.is-transparent {
    background-color: transparent;
    --el-menu-text-color: rgba(255, 255, 255, 0.9);
    --el-menu-item-font-size: 16px;
    
    .container {
      .logo {
        :deep(.el-image) {
          filter: brightness(0) invert(1);
        }
      }
      
      .web-menu {
        .el-menu {
          background-color: transparent;
          
          .el-menu-item {
            color: rgba(255, 255, 255, 0.9);
            background-color: transparent;
            border-bottom: none;
            
            &:hover {
              color: #fff;
              background-color: rgba(255, 255, 255, 0.1);
            }
            
            &.is-active {
              color: #fff;
              border-bottom: 2px solid #fff;
            }
          }
        }
      }
      
      :deep(.auth-popover) {
        color: #fff;
      }
    }
  }
  
  &.is-solid {
    background-color: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    box-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
    --el-menu-text-color: var(--text-primary);
    
    .container {
      padding-top: 8px;
      padding-bottom: 8px;
      
      .web-menu {
        .el-menu {
          background-color: transparent;
          
          .el-menu-item {
            color: var(--text-primary);
            background-color: transparent;
            border-bottom: none;
            
            &:hover {
              color: var(--color-primary);
              background-color: rgba(90, 159, 143, 0.08);
            }
            
            &.is-active {
              color: var(--color-primary);
              border-bottom: 2px solid var(--color-primary);
            }
          }
        }
      }
    }
  }

  .container {
    display: flex;
    align-items: center;
    max-width: 1400px;
    width: 100%;
    padding: 12px 20px;
    transition: padding 0.3s ease;

    .logo {
      height: 50px;
      width: 180px;
      display: flex;
      align-items: center;
    }

    .web-menu {
      margin-left: 40px;

      .el-menu {
        background-color: transparent;
        border-bottom: none;
        --el-menu-item-font-size: 16px;
        --el-menu-item-height: 50px;

        .el-menu-item {
          border-bottom: none;
          background-color: transparent;
          transition: all 0.2s ease;
        }
      }
    }

    .auth-popover {
      margin-left: auto;
    }
  }
}

@media (max-width: 768px) {
  .web-navbar {
    .container {
      .web-menu {
        display: none;
      }
    }
  }
}
</style>
