func mincostTickets(days []int, costs []int) int {
    dps:= make([]int, len(days))
    dur := []int{1, 7, 30}
    for i, v := range days {

        for ii, vv := range dur {
            for j := i; j < len(days); j++ {
                
                if v-1+vv >= days[j] {
                    if i == 0 {
                        if dps[j] == 0 {
                            dps[j] = costs[ii]
                        } else {
                            dps[j] = min(dps[j], costs[ii])
                        }
                    } else {
                        if dps[j] == 0 {
                            dps[j] = dps[i-1]+costs[ii]
                        } else {
                            dps[j] = min(dps[j], dps[i-1]+costs[ii])
                        }
                        
                    }
                }
            }
        }
    }
    return dps[len(days)-1]
}


func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}