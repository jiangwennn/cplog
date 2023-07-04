

<!-- 侧边栏 -->
<template>
  <div class="w-full flex h-screen font-mono">
      <!-- 左侧导航列表 -->
      <div class="h-full flex-none w-20 pt-8 bg-white bg-opacity-5 select-none" style="--wails-draggable:drag">
          <ul class="list-none mt-6" style="--wails-draggable:no-drag">
              <li v-for="item in topNav" class="h-20 w-full">
                  <router-link :to="{name: item.path }" class="block w-full relative top-3" @click="clickNav(item)">
                      <icon class="block w-full text-center">
                        <component :is="item.elem" size="32" :theme="item.theme" :fill="item.fill" :strokeWidth="2" />
                      </icon>
                      <span class="block w-full text-center text-xs" :style="{ color: item.color }">{{ item.name }}</span>
                  </router-link>
              </li>
          </ul>
      </div>
      
      <!-- 右侧内容 -->
      <div class="flex-auto">
          <router-view></router-view>
      </div>
  </div>
</template>

<script setup>
import {
  Like as IconLike,
  Config as IconConfig,
  ViewList as IconListView
} from '@icon-park/vue-next';
import { computed, ref } from 'vue';
import { RouterView, RouterLink } from 'vue-router'

let preset = {
  color: '#333',
  activeColor: '#df1a29',
  theme: 'outline',
  activeTheme: 'multi-color',
  fill: '#333',
  activeFill: ['#df1a29', '#df1a29', '#FFF', '#EA4753']
}


let topNav = ref([
  {
      path: 'History',
      name: '历史',
      theme: preset.activeTheme,
      elem: IconListView,
      fill: preset.activeFill,
      color: preset.activeColor,
  },
  {
      path: 'Like',
      name: '收藏',
      theme: preset.theme,
      elem: IconLike,
      fill: preset.fill,
      color: preset.color,
  },
  {
      path: 'Config',
      name: '设置',
      theme: preset.theme,
      elem: IconConfig,
      fill: preset.fill,
      color: preset.color,
  },
]);

// 点击按钮 切换效果
function clickNav(item) {
  resetNav();
  item.color = preset.activeColor
  item.theme = preset.activeTheme
  item.fill = preset.activeFill
}

// 重置
function resetNav() {
  topNav.value.forEach((val) => {
      val.color = preset.color
      val.fill = preset.fill
      val.theme = preset.theme
  })
}

</script>

<style scoped>
svg{
  display: inline !important;
}
</style>
