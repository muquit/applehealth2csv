# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Applehealth2csv < Formula
  desc "Command line tool to convert Apple Watch Health Data to CSV or JSON files"
  homepage "https://github.com/muquit/applehealth2csv"
  version "1.0.1"
  license "MIT"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/muquit/applehealth2csv/releases/download/v1.0.1/applehealth2csv_1.0.1_mac-64bit.tar.gz"
    sha256 "3e6c97d524d12501d3f62440b72ea948cb40d1ec76497bcfed14fa2ea521b7a5"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/muquit/applehealth2csv/releases/download/v1.0.1/applehealth2csv_1.0.1_linux-64bit.tar.gz"
    sha256 "e3d666ba8bc52ecc803c3168ba6355eec0f107fe39238ba9ecf24bc67b6ea6d6"
  end

  def install
    bin.install "applehealth2csv"
    man1.install "docs/applehealth2csv.1"
  end
end
