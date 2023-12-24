package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type SeedRange struct {
	origin    int
	increment int
}

type AgroRange struct {
	originStart      int
	destinationStart int
	increment        int
}

func (a *AgroRange) inRange(key int) bool {
	return a.originStart <= key && key <= a.originStart+a.increment
}

func (a *AgroRange) getValue(key int) int {
	if a.inRange(key) {
		return key + a.destinationStart - a.originStart
	}

	panic(fmt.Sprintf("The range from %d to %d does not contain the key %d", a.originStart, a.destinationStart, key))
}

type AgroMap struct {
	ranges []AgroRange
}

func (a *AgroMap) getValue(key int) int {
	for _, r := range a.ranges {
		if r.inRange(key) {
			return r.getValue(key)
		}
	}

	return key
}

func main() {
	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	check(err, "Failed to read input file")

	regex, err := regexp.Compile("[0-9]+")
	check(err, "Failed to process regex (you absolute idiot)")
	scanner := bufio.NewScanner(file)

	nextLine := nextLineFromScanner(scanner) // seeds

	// Process seeds
	seedsS := regex.FindAllString(nextLine, -1)
	seeds := make([]SeedRange, len(seedsS)/2)

	// Read seeds in pairs and convert to seed range
	for i := 0; i < len(seedsS); i += 2 {
		originSeed, err := strconv.Atoi(seedsS[i])
		check(err, "Error parsing the origin of the seed range "+seedsS[i])

		seedIncrement, err := strconv.Atoi(seedsS[i+1])
		check(err, "Error parsing the origin of the seed range "+seedsS[i+1])

		seeds[i/2] = SeedRange{origin: originSeed, increment: seedIncrement}
	}

	nextLineFromScanner(scanner) // Empty line

	nextLineFromScanner(scanner) // seed-to-soil map
	seedToSoil := getNextMap(scanner)

	nextLineFromScanner(scanner) // soil-to-fertilizer map
	soilToFertilizer := getNextMap(scanner)

	nextLineFromScanner(scanner) // fertilizer-to-water map
	fertilizerToWater := getNextMap(scanner)

	nextLineFromScanner(scanner) // water-to-light map
	waterToLight := getNextMap(scanner)

	nextLineFromScanner(scanner) // light-to-temperature map
	lightToTemperature := getNextMap(scanner)

	nextLineFromScanner(scanner) // temperature-to-humidity map
	temperatureToHumidity := getNextMap(scanner)

	nextLineFromScanner(scanner) // humidity-to-location map
	humidityToLocation := getNextMap(scanner)

	var lowestLocation int = math.MaxInt
	for _, s := range seeds {
		fmt.Printf("Calculating for range %d to %d\n", s.origin, s.origin+s.increment)
		for i := s.origin; i < s.origin+s.increment; i++ {
			soil := seedToSoil.getValue(i)
			fertilizer := soilToFertilizer.getValue(soil)
			water := fertilizerToWater.getValue(fertilizer)
			light := waterToLight.getValue(water)
			temperature := lightToTemperature.getValue(light)
			humidity := temperatureToHumidity.getValue(temperature)
			location := humidityToLocation.getValue(humidity)

			if location < lowestLocation {
				lowestLocation = location
			}
		}
	}

	fmt.Printf("Lowest location found is %d\n", lowestLocation)
}

func getNextValue(m map[int]int, key int) int {
	if m[key] == 0 {
		return key
	} else {
		return m[key]
	}
}

func getNextMap(s *bufio.Scanner) AgroMap {
	ranges := make([]AgroRange, 0)
	regex, err := regexp.Compile("[0-9]+")
	check(err, "Failed to process regex (you absolute idiot)")

	nextLine := nextLineFromScanner(s)
	for nextLine != "" {
		sts := regex.FindAllString(nextLine, 3)

		destination, err := strconv.Atoi(sts[0])
		check(err, "Failed to parse destination "+sts[0])
		origin, err := strconv.Atoi(sts[1])
		check(err, "Failed to parse origin "+sts[1])
		mapRange, err := strconv.Atoi(sts[2])
		check(err, "Failed to parse range "+sts[2])

		nextRange := AgroRange{originStart: origin, destinationStart: destination, increment: mapRange}
		ranges = append(ranges, nextRange)

		nextLine = nextLineFromScanner(s)
	}

	return AgroMap{ranges: ranges}
}

func nextLineFromScanner(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

func check(e error, message string) {
	if e != nil {
		panic(message)
	}
}
