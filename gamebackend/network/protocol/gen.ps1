# protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto   

# protoc --doc_out=./doc --doc_opt=html,index.html proto/*.proto

protoc --validate_out="lang=go:./gen" --go_out=./gen/ --doc_out=./doc --doc_opt=html,index.html proto/*.proto
