package main
  
import (
    "fmt"
    "math"
)

var mod = math.Pow(10, 9);

type hPaths struct {
    From    int
    To      int
}

var path []hPaths

//Main Function
func main() {

    var cost int

    cost = o(2, 5, 1, 3, 5)
    fmt.Println("Expected output is 60 and actually get cost of E(2,5,1,3,5) is",cost)

    cost = o(3, 20, 4, 9, 17)
    fmt.Println("Expected output is 2358 and actually get cost of E(3, 20, 4, 9, 17) is",cost)
    
    // Get the summation cost
    SummationCost()
}

//Move Disks
func TowerOfHanoi(n int, from int, to int, aux int) {

    if n == 1 {
        if (len(path) == 0) {
            if to < aux {
                path = append(path, hPaths{to, from})
            } else {
                path = append(path, hPaths{aux, from})
            }
        } else {
            var last = path[len(path) - 1]
            path = append(path, hPaths{last.To, from})
        }
        path = append(path, hPaths{from, to})
        return;
    }

    TowerOfHanoi(n-1, from, aux, to)

    if (len(path) == 0) {
        path = append(path, hPaths{aux, from})
    } else {
        var last = path[len(path) - 1]
        path = append(path, hPaths{last.To, from})
    }

    path = append(path, hPaths{from, to})

    TowerOfHanoi(n-1, aux, to, from)
}

//Move Cost
func moveCost(tiles int, i int, j int) int {
    
    var cost int
    if i < j {
        cost = int(math.Pow((float64(j) - 1), 2) - math.Pow((float64(i) - 1), 2))
    } else {
        cost = int(math.Pow((float64(tiles) - float64(j)), 2) - math.Pow((float64(tiles) - float64(i)), 2))
    }
    return int(cost % int(mod))
}

// Output Function
func o(nDisk int, nTiles int, towerA int, towerB int, towerC int) int {
    
    path = []hPaths{}

    TowerOfHanoi(nDisk, towerA, towerC, towerB)

    var pathCost = 0

    for _, element := range path {
        pathCost += moveCost(nTiles, element.From, element.To)
    }

    pathCost = pathCost % int(mod)

    return pathCost

}

// Summation cost
func SummationCost() {
    var summationCost = 0
    var i int

    for i <= 100 {
        var nDisk = i % int(mod)
        var nTiles = (nDisk * 10) % int(mod)
        var towerA = (nDisk * 3) % int(mod)
        var towerB = (nDisk * 6) % int(mod)
        var towerC = (nDisk * 9) % int(mod)
        path = []hPaths{}

        TowerOfHanoi(nDisk, towerA, towerC, towerB)

        var pathCost = 0

        for _, element := range path {
            pathCost += moveCost(nTiles, element.From, element.To)
        }
        summationCost += pathCost
        summationCost = summationCost % int(mod)
        i = i + 1
    }

    fmt.Println("Summation Cost is ",summationCost)
}