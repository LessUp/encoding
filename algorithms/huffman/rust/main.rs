use std::io;

use compresskit_codec::cli;
use compresskit_codec::codec::{decode_buffer, encode_buffer};

struct HuffmanProcessor;

impl cli::FileProcessor for HuffmanProcessor {
    fn encode_file(&self, input_path: &str, output_path: &str) -> io::Result<()> {
        let input = std::fs::read(input_path).map_err(|e| {
            io::Error::new(
                e.kind(),
                format!("cannot open input file for reading: {input_path}: {e}"),
            )
        })?;
        let mut encoder = huffman::new_encoder();
        let encoded = encode_buffer(&mut encoder, &input)
            .map_err(|e| io::Error::new(io::ErrorKind::Other, e.to_string()))?;
        std::fs::write(output_path, encoded).map_err(|e| {
            io::Error::new(
                e.kind(),
                format!("cannot open output file for writing: {output_path}: {e}"),
            )
        })
    }

    fn decode_file(&self, input_path: &str, output_path: &str) -> io::Result<()> {
        let input = std::fs::read(input_path).map_err(|e| {
            io::Error::new(
                e.kind(),
                format!("cannot open input file for reading: {input_path}: {e}"),
            )
        })?;
        let mut decoder = huffman::new_decoder();
        let decoded = decode_buffer(&mut decoder, &input)
            .map_err(|e| io::Error::new(io::ErrorKind::Other, e.to_string()))?;
        std::fs::write(output_path, decoded).map_err(|e| {
            io::Error::new(
                e.kind(),
                format!("cannot open output file for writing: {output_path}: {e}"),
            )
        })
    }
}

fn main() {
    cli::run("huffman", &HuffmanProcessor);
}
