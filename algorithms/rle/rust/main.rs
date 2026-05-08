use std::io;

use compresskit_codec::cli;
use compresskit_codec::codec::{decode_buffer, encode_buffer};

struct RLEProcessor;

impl cli::FileProcessor for RLEProcessor {
    fn encode_file(&self, input_path: &str, output_path: &str) -> io::Result<()> {
        let input = std::fs::read(input_path).map_err(|e| {
            io::Error::new(
                e.kind(),
                format!("cannot open input file for reading: {input_path}: {e}"),
            )
        })?;
        let mut encoder = rle::new_encoder();
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
        let mut decoder = rle::new_decoder();
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
    cli::run("rle", &RLEProcessor);
}
