<template>
  <div class="hero-banner">
    <div class="banner-background" :style="{ backgroundImage: `url(${currentImage})` }">
      <div class="banner-overlay"></div>
    </div>
    <div class="banner-content">
      <div class="banner-text">
        <h1 class="site-title">{{ websiteStore.state.websiteInfo.name || '我的博客' }}</h1>
        <p class="site-subtitle">{{ websiteStore.state.websiteInfo.description || '记录生活，分享技术' }}</p>
        <div class="scroll-hint" @click="scrollToContent">
          <el-icon class="scroll-icon">
            <ArrowDown />
          </el-icon>
        </div>
      </div>
    </div>
    <!-- 图片切换指示器 -->
    <div class="banner-indicators">
      <span 
        v-for="(item, index) in imgList" 
        :key="index"
        :class="['indicator', { active: currentIndex === index }]"
        @click="setCurrentImage(index)"
      ></span>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, onMounted, onUnmounted, computed} from "vue";
import {websiteCarousel} from "@/api/website";
import {useWebsiteStore} from "@/stores/website";
import {ArrowDown} from '@element-plus/icons-vue';

const websiteStore = useWebsiteStore()

const imgList = ref<string[]>([
  '/image/carousel_1.png',
  '/image/carousel_2.png',
  '/image/carousel_3.png',
  '/image/carousel_4.png',
])

const currentIndex = ref(0)
let timer: number | null = null

const currentImage = computed(() => imgList.value[currentIndex.value])

const setCurrentImage = (index: number) => {
  currentIndex.value = index
}

const nextImage = () => {
  currentIndex.value = (currentIndex.value + 1) % imgList.value.length
}

const startAutoPlay = () => {
  timer = window.setInterval(() => {
    nextImage()
  }, 6000)
}

const scrollToContent = () => {
  const content = document.querySelector('.main-content')
  if (content) {
    content.scrollIntoView({ behavior: 'smooth' })
  }
}

const getWebsiteCarousel = async () => {
  const res = await websiteCarousel()
  if (res.code === 0 && res.data.length !== 0) {
    imgList.value = res.data
  }
}

onMounted(() => {
  getWebsiteCarousel()
  startAutoPlay()
})

onUnmounted(() => {
  if (timer) {
    clearInterval(timer)
  }
})
</script>

<style scoped lang="scss">
.hero-banner {
  position: relative;
  width: 100%;
  height: 100vh;
  min-height: 600px;
  overflow: hidden;
}

.banner-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  transition: background-image 1s ease-in-out;
}

.banner-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    to bottom,
    rgba(0, 0, 0, 0.1) 0%,
    rgba(0, 0, 0, 0.3) 50%,
    rgba(0, 0, 0, 0.5) 100%
  );
}

.banner-content {
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  padding: 0 20px;
}

.banner-text {
  text-align: center;
  color: #fff;
  animation: fadeInUp 1s ease-out;
}

.site-title {
  font-size: 4rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
  font-family: 'Brush Script MT', 'Lucida Handwriting', 'Courier New', cursive;
  color: #ffffff;
  text-shadow:
    0 0 8px rgba(255, 255, 255, 0.4),
    0 0 16px rgba(255, 255, 255, 0.2),
    2px 2px 10px rgba(0, 0, 0, 0.3);
  position: relative;
  display: block;
  animation: titleGlow 4s ease-in-out infinite alternate;
  transform: rotate(-2deg);
  
  &::after {
    content: '';
    position: absolute;
    bottom: -8px;
    left: 0;
    width: 100%;
    height: 2px;
    background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.6), transparent);
    border-radius: 50%;
  }
}

.site-subtitle {
  font-size: 1.6rem;
  font-weight: 400;
  opacity: 0.9;
  margin-bottom: 3rem;
  font-family: 'Lucida Handwriting', 'Courier New', cursive;
  color: #f0f8ff;
  text-shadow:
    0 0 5px rgba(255, 255, 255, 0.3),
    1px 1px 6px rgba(0, 0, 0, 0.2);
  position: relative;
  display: block;
  animation: subtitleFloat 5s ease-in-out infinite;
  transform: rotate(1deg);
  margin-top: 0.5rem;
}

.scroll-hint {
  cursor: pointer;
  animation: bounce 2s infinite;
  transition: opacity 0.3s;
  
  &:hover {
    opacity: 0.7;
  }
}

.scroll-icon {
  font-size: 2.5rem;
  color: #fff;
}

.banner-indicators {
  position: absolute;
  bottom: 40px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 3;
  display: flex;
  gap: 12px;
}

.indicator {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background-color: rgba(255, 255, 255, 0.4);
  cursor: pointer;
  transition: all 0.3s ease;
  
  &:hover {
    background-color: rgba(255, 255, 255, 0.7);
  }
  
  &.active {
    background-color: #fff;
    transform: scale(1.2);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-15px);
  }
  60% {
    transform: translateY(-8px);
  }
}

@keyframes titleGlow {
  0% {
    filter: brightness(1) drop-shadow(0 0 3px rgba(255, 255, 255, 0.3));
    transform: rotate(-2deg) scale(1);
  }
  50% {
    filter: brightness(1.05) drop-shadow(0 0 8px rgba(255, 255, 255, 0.5));
    transform: rotate(-1.5deg) scale(1.02);
  }
  100% {
    filter: brightness(1.1) drop-shadow(0 0 12px rgba(255, 255, 255, 0.7));
    transform: rotate(-2deg) scale(1);
  }
}

@keyframes subtitleFloat {
  0%, 100% {
    transform: rotate(1deg) translateY(0);
  }
  50% {
    transform: rotate(0.5deg) translateY(-3px);
  }
}

@media (max-width: 768px) {
  .site-title {
    font-size: 2.8rem;
    transform: rotate(-1deg);
    
    &::after {
      bottom: -6px;
      height: 1.5px;
    }
  }
  
  .site-subtitle {
    font-size: 1.4rem;
    transform: rotate(0.5deg);
  }
  
  .hero-banner {
    min-height: 500px;
  }
}
</style>
