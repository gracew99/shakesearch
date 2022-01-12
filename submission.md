Web app is live here: https://limitless-headland-54771.herokuapp.com/
Video walkthrough: https://streamable.com/b2ytwi

Features Implemented:
1. Highlighting found keywords
2. Multi word search (return results matching whole query as well as individual words. ie: "hello shakespeare fans" would return matches for "hello shakespeare fans" and "hello" "shakespeare" and "fans")
3. Results returned in the order they appear in the text
4. UI improvement (switched to React): add background, paginated table for displaying results, search bar, home page image, separate pages for home page/results/read more, etc
5. Case-insensitive search
6. Read More page- links on results page point to a page where the user can read a larger surrounding context and also go back/forth between pages. The correspnding result hit is highlighted on the read more page.
7. Merge overlapping search hits to reduce number of entries returned (ie: if two words are found in close proximity, return both in one result instead of two separate results)
8. Instead of cutting of returned excerpts in the middle of the word, returned windows are full lines. 
9. Display no results found message if applicable.
10. Handle edge case on results page: matches that appeear at the very beginning/end of the document originally are out of bounds. 

Future Changes:
1. Loading icon while fetching results (for queries of "a" for example, the page intially displays 0 results).
2. Return more detailed information, ie: the particular play the search hit is from. Also allow more granular search ie: search a particular work
3. Scroll back to the top when navigating the table. Also reset the table to page 1 for subsequent queries. Also the user currently has to clear their current search before making a new one; perhaps clear the search bar between queries. 
4. Make the search more robust, for example: handle typos or punctuation (ie: searching "hamlets" returns results including "hamlet's")
5. Improve multi word search: right now, it just returns results matching the whole search string or each of the words individually but it should match subsets of the search string too (ie: "hello there world" should return results for "hello there" and "there world" too). Also, the results should be ordered by best match (ie: matches for "hello there world" should appear first).
6. Instead of having read more links for each individual index of a match, have each entry in the table map to one read more link. Then on the read more page, all words found in that particular entry should be highlighted.
