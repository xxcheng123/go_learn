package t_runtime

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/8 15:35
 */

// // lockRankStruct is embedded in mutex, but is empty when staticklockranking is
// disabled (the default)
type lockRankStruct struct {
}

type lockRank int
