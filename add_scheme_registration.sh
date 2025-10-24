#!/bin/bash

# Find all files with *List types
for file in $(find apis/ -name "*types.go" -exec grep -l "type.*List struct" {} \;); do
    echo "Processing $file"
    
    # Find the type name (e.g., LoadBalancerList -> LoadBalancer)
    list_type=$(grep "type.*List struct" "$file" | sed 's/type \([A-Z][a-zA-Z]*\)List struct/\1/')
    
    if [ -n "$list_type" ]; then
        echo "Found type: $list_type"
        
        # Check if init function already exists
        if ! grep -q "SchemeBuilder.Register.*${list_type}" "$file"; then
            # Find the end of the List type definition
            line_num=$(grep -n "type ${list_type}List struct" "$file" | cut -d: -f1)
            if [ -n "$line_num" ]; then
                # Find the closing brace of the struct
                end_line=$(tail -n +$line_num "$file" | grep -n "^}" | head -1 | cut -d: -f1)
                if [ -n "$end_line" ]; then
                    actual_end=$((line_num + end_line - 1))
                    echo "Adding init function after line $actual_end"
                    
                    # Create temp file
                    sed "${actual_end}a\\
\\
func init() {\\
\tSchemeBuilder.Register(\&${list_type}{}, \&${list_type}List{})\\
}" "$file" > "${file}.tmp" && mv "${file}.tmp" "$file"
                fi
            fi
        else
            echo "Init function already exists for $list_type"
        fi
    fi
done
