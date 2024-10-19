package geohashgo

import (
	"math"

	"github.com/bluele/gcache"
)

const (
	MINLAT float64 = -90.0
	MAXLAT float64 = 90.0
	MINLON float64 = -180.0
	MAXLON float64 = 180.0
)

var (
	base32Table = []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '2', '3', '4', '5', '6', '7',
	}
)

type GeoHash interface {
	EncodeBits(numbits int, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon float64) []byte
	EncodeBase32(numbits int, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon float64) []byte
	EncodeBitsWithPerturbationDistribution(numbits int, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon float64) []byte
	EncodeBase32WithPerturbationDistribution(numbits int, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon float64) []byte
	DecodeBits(hashbits []byte) (resLat, resLon float64)
	DecodeBase32(hashbits []byte) (resLat, resLon float64)
}

func NewGeoHash() GeoHash {
	return &geoHash{}
}

type geoHash struct {
	cache *gcache.LRUCache
}

func (g *geoHash) EncodeBits(numbits int, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon float64) []byte {
	key := targetLat*1000 + targetLon
	if g.cache.Has(key) {
		val, err := g.cache.Get(key)
		if err != nil {
			goto defaultCal
		}
		ret, _ := val.([]byte)
		return ret
	}
defaultCal:
	byteNums := int(math.Ceil(float64(numbits) / 8))
	latBytes := make([]byte, byteNums)
	lonBytes := make([]byte, byteNums)
	for i := 0; i < numbits; i++ {
		latMid := (ceilLat + floorLat) * 0.5
		lonMid := (ceilLon + floorLon) * 0.5
		if targetLat >= latMid {
			floorLat = targetLat
		} else {
			ceilLat = targetLat
		}
		if targetLon >= lonMid {
			floorLon = targetLon
		} else {
			ceilLon = targetLon
		}
	}
	ret := append(latBytes, lonBytes...)
	g.cache.Set(key, ret)
	return ret
}

func (g *geoHash) EncodeBase32(numbits int, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon float64) []byte {
	key := targetLat*10000 + targetLon
	if g.cache.Has(key) {
		val, err := g.cache.Get(key)
		if err != nil {
			goto defaultCal
		}
		ret, _ := val.([]byte)
		return ret
	}
defaultCal:
	geoHashBits := g.EncodeBits(numbits, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon)
	ret := encodeBase32(geoHashBits)
	g.cache.Set(key, ret)
	return ret
}

func encodeBase32(bits []byte) []byte {
	var ret []byte
	return ret
}

func (g *geoHash) EncodeBitsWithPerturbationDistribution(numbits int, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon float64) []byte
func (g *geoHash) EncodeBase32WithPerturbationDistribution(numbits int, ceilLat, floorLat, ceilLon, floorLon, targetLat, targetLon float64) []byte
func (g *geoHash) DecodeBits(hashbits []byte) (resLat, resLon float64)
func (g *geoHash) DecodeBase32(hashbits []byte) (resLat, resLon float64)
