func reversePairs(nums []int) int {
    var ans int
    mergeCount(nums, 0, len(nums)-1, &ans)
    return ans
}

func mergeCount(nums []int, start, end int, count *int) []int {
    if start > end {
        return nil
    }
    if start == end {
        return []int{nums[start]}
    }
    mid := start + (end - start)/2
    l := mergeCount(nums, start, mid, count)
    r := mergeCount(nums, mid+1, end, count)
    return merge(l, r, count)
}

func merge(l, r []int, count *int) []int {
    var i, j int
    ret := make([]int, 0)
    for i < len(l) && j < len(r) {
        if l[i] > r[j] {
            *count += len(l)-i // len(l)-1-i+1
            ret = append(ret, r[j])
            j++
        } else {
            ret = append(ret, l[i])
            i++
        }
    }
    for i < len(l) {
        ret = append(ret, l[i])
        i++
    }
    for j < len(r) {
        ret = append(ret, r[j])
        j++
    }
    return ret
}
