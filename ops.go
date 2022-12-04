package lib1n

// Return new dataset, where old dataset matched
// Warning: Concurrent function
func Find(ds DataSet, pattern string) DataSet {
	type result struct {
		key string
		val []string
	}

	var chResult = make(chan result, len(ds))

	// For each key/value pair, try concurrently
	// to match the pattern
	for key, val := range ds {
		go func(pattern string, key string, vals []string) {
			for _, val := range vals {
				// If pattern matched:
				// Send trough channel, and return
				if val == pattern {
					chResult <- result{key: key, val: vals}
					return
				}
			}
			// If no pattern matched, send empty
			chResult <- result{key: "", val: nil}
		}(pattern, key, val)
	}

	nds := make(DataSet, len(ds))

	// Collect results, and return new dataset
	for i := 0; i < len(ds); i++ {
		res := <-chResult
		if res.val != nil {
			nds[res.key] = res.val
		}
	}

	return nds
}

// Return new dataset, where old values are replaced
// Warning: Concurrent function
func Replace(ds DataSet, pattern string, replace string) DataSet {
	type result struct {
		key string
		val []string
	}

	var chResult = make(chan result, len(ds))

	// For each key/value pair, try concurrently
	// to match and replace pattern
	for key, val := range ds {
		go func(pattern string, replace string, key string, vals []string) {
			for i, val := range vals {
				// If pattern matched, replace it
				if val == pattern {
					vals[i] = replace
				}
			}
			chResult <- result{key: key, val: vals}
		}(pattern, replace, key, val)
	}

	nds := make(DataSet, len(ds))

	// Collect results, and return new dataset
	for i := 0; i < len(ds); i++ {
		res := <-chResult
		nds[res.key] = res.val
	}

	return nds
}

// Execute function on each dataframe
func Map(ds DataSet, fn func(df []string) []string) DataSet {
	for key, val := range ds {
		ds[key] = fn(val)
	}
	return ds
}
