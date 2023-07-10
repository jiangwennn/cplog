<template>
    <div class="flex h-full">
        <div class="basis-80 flex-none h-full flex select-none">
            <div class="w-full h-full bg-white box-border border-r border-neutral-200 flex flex-col">
                <div class="basis-16 flex-none w-full bg-white flex items-center box-border border-b border-neutral-200">
                    <div class="basis-5/6 h-8 ml-2 flex items-center relative">
                        <input type="text" placeholder="搜索历史记录"
                            class="w-full h-full rounded-md indent-7 text-xs bg-zinc-100 focus:outline-none focus:ring-0 focus:border focus:border-[#df1a29]" />
                        <IconSearch class="absolute ml-2 top-[3px]" size="14" stroke-width="1" />
                    </div>
                    <div class="basis-1/6 text-center">
                        <IconPlus class="bg-zinc-100 block h-full w-[26px] mx-auto rounded-sm active:bg-zinc-50"
                            theme="outline" size="24" fill="#333" stroke-width="2" />
                    </div>
                </div>
                <div class="flex-auto overflow-y-auto cursor-default">
                    <div v-for="item in historyList" @click="details = item"
                        class="w-full h-24 border-b border-neutral-100 text-xs active:bg-neutral-50">
                        <div class="relative top-4 w-11/12 mx-auto h-8 text-neutral-600 box-border flex">
                            <span class="basis-10/12 line-clamp-2">
                                {{ item.Content }}
                            </span>
                            <icon class="basis-2/12 w-full ">
                                <IconCopy class="w-full h-full text-center flex items-center justify-center" size="16"
                                    stroke-width="1" />
                            </icon>
                        </div>
                        <div class="relative top-8 w-11/12 mx-auto h-4 flex">
                            <span class="w-full h-full basis-3/4 text-neutral-400 text-xs"> {{ item.Time }}</span>
                            <icon class="basis-1/4 flex justify-end">
                            </icon>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- 右侧详情 -->
        <div class="h-full flex-auto">
            <div v-if="details.Id !== undefined" class="w-full h-full bg-white">
                <div class="w-full h-16 box-border border-b border-neutral-200 cursor-default flex justify-end select-none"
                    style="--wails-draggable:drag">
                    <div class="flex-auto h-full text-center flex items-center">
                        <span class="flex-none ml-4">详情</span>
                    </div>
                    <icon class="basis-10 h-full flex items-center justify-start">
                        <IconCopy class="text-center" size="20" stroke-width="1" />
                    </icon>
                    <icon class="basis-10 h-full flex items-center justify-start">
                        <IconLike class="text-center" size="20" stroke-width="1" />
                    </icon>
                    <icon class="basis-10 h-full flex items-center justify-start">
                        <IconDelete class="text-center" size="20" stroke-width="1" />
                    </icon>
                </div>
                <div class="w-full h-full">
                    <div class="border border-slate-50 w-full h-full bg-slate-50 p-4">
                        {{ details.Content }}
                    </div>
                </div>
            </div>

            <!-- 主页介绍 -->
            <About v-else style="--wails-draggable:drag" class="bg-gray-100 h-full mx-auto w-full select-none"></About>
        </div>
    </div>
</template>

<script setup>
import About from '@/components/About.vue';
import { useStore } from 'vuex';
import { computed, ref } from 'vue';

import { LogPrint } from '_wailsjs/runtime/runtime.js'
import {
    Plus as IconPlus,
    Search as IconSearch,
    Like as IconLike,
    Delete as IconDelete,
    Copy as IconCopy,
} from '@icon-park/vue-next';


const store = useStore();


store.dispatch('history/start')

let historyList = computed(() => store.state.history.list);

let details = ref({})



</script>

<style></style>